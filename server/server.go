package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go_gorm_graphql/resolver"
)

func Init() {
	r := router()
	r.Run(":3000")
}

func router() *gin.Engine {
	r := gin.Default()
	r.POST("/graphql", graphqlHandler())
	r.GET("/query", playgroundHandler())

	return r
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	config := resolver.Config{
		Resolvers: resolver.New(),
	}
	h := handler.NewDefaultServer(resolver.NewExecutableSchema(config))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
