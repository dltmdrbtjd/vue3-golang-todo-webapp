package main

import (
	"fmt"

	"github.com/Convenience-Tools/convenience-server/configs"
	"github.com/Convenience-Tools/convenience-server/pkg/middleware"
	"github.com/Convenience-Tools/convenience-server/pkg/routes"
	"github.com/Convenience-Tools/convenience-server/pkg/schema"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.JSONMiddleware())

	configs.ConnectDB()
	routes.TodoRoutes(r)
	r.POST("/query", GraphQLHandler())

	if err := r.Run(":8081"); err != nil {
		fmt.Printf("failed server start, err%v\n", err)
	}
}

func GraphQLHandler() gin.HandlerFunc {
	h := relay.Handler{
		Schema: schema.Schema,
	}

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
