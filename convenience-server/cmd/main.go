package main

import (
	"fmt"

	"github.com/Convenience-Tools/convenience-server/configs"
	"github.com/Convenience-Tools/convenience-server/pkg/middleware"
	"github.com/Convenience-Tools/convenience-server/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.JSONMiddleware())

	configs.ConnectDB()
	routes.TodoRoutes(r)
	routes.GraphQLRouter(r)

	if err := r.Run(":8081"); err != nil {
		fmt.Printf("failed server start, err%v\n", err)
	}
}
