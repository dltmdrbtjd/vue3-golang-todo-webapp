package todo

import (
	"github.com/Convenience-Tools/convenience-server/pkg/controllers/todo"
	"github.com/gin-gonic/gin"
)

func Register(parentGroup *gin.RouterGroup, controller todo.Controller) {
	if controller == nil {
		controller = todo.CreateController()
	}
	todoGroup := parentGroup.Group("/todo")
	todoGroup.Use()
	{
		todoGroup.POST("", controller.CreateTodoItem)
		todoGroup.GET("/:username", controller.GetTodoList)
		todoGroup.DELETE("/:todoId", controller.DeleteTodoItem)
	}
}
