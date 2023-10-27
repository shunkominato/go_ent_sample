package main

import (
	"log"
	"net/http"
	"os"

	"go-gql-sample/app/ent"
	Dataloader "go-gql-sample/app/internal/dataloader"
	"go-gql-sample/app/internal/infrastructure/server/graph"
	"go-gql-sample/app/internal/infrastructure/server/graph/resolver"
	"go-gql-sample/app/pkg/config"
	"go-gql-sample/app/pkg/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func graphqlHandler(client *ent.Client) gin.HandlerFunc {	
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolver.Resolver{
			Client: client,
		},
	}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	godotenv.Load(".env")

	config.SetConfig()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()

	db, _ := db.NewDatabase()	
	client := db.EntClient()
	defer db.Close()

	loaders := Dataloader.NewLoaders(client)
	r.Use(Dataloader.Middleware(loaders))

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.POST("/query", graphqlHandler(client))
	r.GET("/", playgroundHandler())

	r.Run()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
