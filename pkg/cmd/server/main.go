package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/backend/pkg/auth"
	"github.com/scorify/backend/pkg/cache"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/data"
	"github.com/scorify/backend/pkg/engine"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/inject"
	"github.com/scorify/backend/pkg/ent/user"
	"github.com/scorify/backend/pkg/graph"
	"github.com/scorify/backend/pkg/graph/directives"
	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/scorify/backend/pkg/grpc/server"
	"github.com/scorify/backend/pkg/structs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "Run the server",
	Long:    "Run the server",
	Aliases: []string{"s", "serve"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init()
	},

	Run: run,
}

func graphqlHandler(entClient *ent.Client, redisClient *redis.Client, engineClient *engine.Client, scoreTaskChan chan *proto.GetScoreTaskResponse, scoreTaskReponseChan chan *proto.SubmitScoreTaskRequest) gin.HandlerFunc {
	conf := graph.Config{
		Resolvers: &graph.Resolver{
			Ent:                  entClient,
			Redis:                redisClient,
			Engine:               engineClient,
			ScoreTaskChan:        scoreTaskChan,
			ScoreTaskReponseChan: scoreTaskReponseChan,
		},
	}

	conf.Directives.IsAuthenticated = directives.IsAuthenticated
	conf.Directives.HasRole = directives.HasRole

	h := handler.New(
		graph.NewExecutableSchema(
			conf,
		),
	)

	h.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	if gin.IsDebugging() {
		h.Use(extension.Introspection{})
	}

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func injectFileHandler(entClient *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		entUser, err := auth.Parse(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		fileType := c.Param("fileType")
		parentID := c.Param("parentID")
		fileID := c.Param("fileID")
		fileName := c.Param("filename")

		if fileType != string(structs.FileTypeInject) && fileType != string(structs.FileTypeSubmission) {
			c.JSON(http.StatusNotFound, gin.H{"error": "file type not found"})
			return
		}

		parentUUID, err := uuid.Parse(parentID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid parent id"})
			return
		}

		fileUUID, err := uuid.Parse(fileID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid file id"})
			return
		}

		var entInject *ent.Inject

		if entUser.Role == user.RoleAdmin {
			// Admins can access all files
			entInject, err = entClient.Inject.Query().
				Where(
					inject.ID(parentUUID),
				).
				Only(c)
		} else {
			// Users can only access files that are currently active
			now := time.Now()
			entInject, err = entClient.Inject.Query().
				Where(
					inject.ID(parentUUID),
					inject.StartTimeLTE(now),
					inject.EndTimeGTE(now),
				).
				Only(c)
		}
		if err != nil {
			logrus.WithError(err).Error("failed to get inject")
			c.JSON(http.StatusNotFound, gin.H{"error": "inject not found"})
			return
		}

		var file *structs.File

		for _, f := range entInject.Files {
			if f.ID == fileUUID && f.Name == fileName {
				file = &structs.File{
					ID:   f.ID,
					Name: f.Name,
				}
				break
			}
		}

		if file == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
			return
		}

		filePath, err := file.FilePath(structs.FileType(fileType), parentUUID)
		if err != nil {
			logrus.WithError(err).Error("failed to get file path")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.File(filePath)
	}
}

func startWebServer(wg *sync.WaitGroup, entClient *ent.Client, redisClient *redis.Client, engineClient *engine.Client, scoreTaskChan chan *proto.GetScoreTaskResponse, scoreTaskReponseChan chan *proto.SubmitScoreTaskRequest) {
	defer wg.Done()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(auth.JWTMiddleware(entClient))

	err := router.SetTrustedProxies(nil)
	if err != nil {
		logrus.WithError(err).Fatal("failed to set trusted proxies")
	}

	cors_urls := []string{
		fmt.Sprintf("http://%s:%d", config.Domain, config.Port),
		fmt.Sprintf("https://%s:%d", config.Domain, config.Port),
		fmt.Sprintf("http://%s:3000", config.Domain),
		fmt.Sprintf("https://%s:3000", config.Domain),
		fmt.Sprintf("http://%s:5173", config.Domain),
		fmt.Sprintf("https://%s:5173", config.Domain),
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     cors_urls,
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	router.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	router.POST("/query", graphqlHandler(entClient, redisClient, engineClient, scoreTaskChan, scoreTaskReponseChan))
	router.GET("/query", graphqlHandler(entClient, redisClient, engineClient, scoreTaskChan, scoreTaskReponseChan))
	router.GET("/api/files/inject/:parentID/:fileID/:fileName", injectFileHandler(entClient))

	logrus.Printf("Starting web server on http://%s:%d", config.Domain, config.Port)

	err = router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logrus.WithError(err).Fatal("failed to start server")
	} else {
		logrus.Info("Server stopped")
	}
}

func startGRPCServer(wg *sync.WaitGroup, scoreTaskChan chan *proto.GetScoreTaskResponse, scoreTaskReponseChan chan *proto.SubmitScoreTaskRequest) {
	defer wg.Done()

	server.Serve(
		context.Background(),
		scoreTaskChan,
		scoreTaskReponseChan,
	)
}

// serverRun runs the server
func run(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()

	entClient, err := data.NewClient(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create ent client")
	}

	scoreTaskChan := make(chan *proto.GetScoreTaskResponse)
	scoreTaskReponseChan := make(chan *proto.SubmitScoreTaskRequest)
	defer close(scoreTaskChan)
	defer close(scoreTaskReponseChan)

	redisClient := cache.NewRedisClient()

	engineClient := engine.NewEngine(ctx, entClient, redisClient, scoreTaskChan, scoreTaskReponseChan)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go startWebServer(wg, entClient, redisClient, engineClient, scoreTaskChan, scoreTaskReponseChan)
	go startGRPCServer(wg, scoreTaskChan, scoreTaskReponseChan)

	wg.Wait()
}
