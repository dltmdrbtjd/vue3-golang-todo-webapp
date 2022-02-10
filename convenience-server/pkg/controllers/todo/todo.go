package todo

import (
	"net/http"

	todoModel "github.com/Convenience-Tools/convenience-server/pkg/models/todo"
	"github.com/Convenience-Tools/convenience-server/pkg/response"
	"github.com/Convenience-Tools/convenience-server/pkg/services/todo"
	"github.com/gin-gonic/gin"
)

type controller struct {
	todoService todo.Service
}

type Controller interface {
	CreateTodoItem(c *gin.Context)
}

func NewController(todoService todo.Service) *controller {
	return &controller{
		todoService: todoService,
	}
}

func CreateController() *controller {
	return NewController(todo.CreateService())
}

func (ctrl *controller) CreateTodoItem(c *gin.Context) {
	var todoItem *todoModel.Todo
	if err := c.BindJSON(&todoItem); err != nil {
		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	err := ctrl.todoService.CreateTodo(todoItem.Title, todoItem.Status, todoItem.Content, todoItem.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusCreated, response.TodoResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{}})
}
