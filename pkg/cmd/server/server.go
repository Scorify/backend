package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/scorify/backend/pkg/auth"
	"github.com/scorify/backend/pkg/cache"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/data"
	"github.com/scorify/backend/pkg/graph"
	"github.com/scorify/backend/pkg/graph/directives"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "Run the server",
	Long:    "Run the server",
	Aliases: []string{"s", "serve"},

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

// serverRun runs the server
func run(cmd *cobra.Command, args []string) {
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

	logrus.Printf("Starting server on http://%s:%d", config.Domain, config.Port)

	err = router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logrus.WithError(err).Fatal("failed to start server")
	} else {
		logrus.Info("Server stopped")
	}
}
