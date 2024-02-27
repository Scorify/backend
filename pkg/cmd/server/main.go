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
	"github.com/scorify/backend/pkg/auth"
	"github.com/scorify/backend/pkg/cache"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/data"
	"github.com/scorify/backend/pkg/graph"
	"github.com/scorify/backend/pkg/graph/directives"
	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/scorify/backend/pkg/grpc/server"
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
		data.Init()
		cache.Init()
	},

	Run: run,
}

func graphqlHandler() gin.HandlerFunc {
	conf := graph.Config{
		Resolvers: &graph.Resolver{
			Ent:   data.Client,
			Redis: cache.Client,
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

func startWebServer(wg *sync.WaitGroup) {
	defer wg.Done()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(auth.JWTMiddleware)

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
	router.POST("/query", graphqlHandler())
	router.GET("/query", graphqlHandler())

	logrus.Printf("Starting web server on http://%s:%d", config.Domain, config.Port)

	err = router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logrus.WithError(err).Fatal("failed to start server")
	} else {
		logrus.Info("Server stopped")
	}
}

func startGRPCServer(wg *sync.WaitGroup) {
	scoreTaskChan := make(chan *proto.GetScoreTaskResponse)
	scoreTaskReponseChan := make(chan *proto.SubmitScoreTaskRequest)

	defer wg.Done()
	defer close(scoreTaskChan)
	defer close(scoreTaskReponseChan)

	go func() {
		i := 0
		for {
			time.Sleep(2 * time.Second)
			scoreTaskChan <- &proto.GetScoreTaskResponse{
				StatusId:   uuid.New().String(),
				SourceName: fmt.Sprintf("source-%d", i),
				Config:     "{}",
			}
			i++
		}
	}()

	go func() {
		for range scoreTaskReponseChan {
		}
	}()

	server.Serve(
		context.Background(),
		scoreTaskChan,
		scoreTaskReponseChan,
	)
}

// serverRun runs the server
func run(cmd *cobra.Command, args []string) {
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go startWebServer(wg)
	go startGRPCServer(wg)

	wg.Wait()
}
