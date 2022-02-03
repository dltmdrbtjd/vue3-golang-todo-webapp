package routes

import (
	"github.com/Convenience-Tools/convenience-server/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(route *gin.Engine) {
	route.POST("/todo", controllers.CreateTodo)
	route.GET("/todo/:userName", controllers.GetTodoList)
	route.DELETE("/todo/:objId", controllers.DeleteTodo)
	route.PUT("/todo/:objId", controllers.EditTodo)
	route.PUT("/todo/status/:objId", controllers.TodoStatusChange)
}
