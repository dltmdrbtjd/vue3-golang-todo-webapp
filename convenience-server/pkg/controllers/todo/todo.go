package todo

import (
	"net/http"

	todoModel "github.com/Convenience-Tools/convenience-server/pkg/models/todo"
	"github.com/Convenience-Tools/convenience-server/pkg/response"
	"github.com/Convenience-Tools/convenience-server/pkg/services/todo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type controller struct {
	todoService todo.Service
}

type Controller interface {
	CreateTodoItem(c *gin.Context)
	GetTodoList(c *gin.Context)
	DeleteTodoItem(c *gin.Context)
	EditTodoItem(c *gin.Context)
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

	err := ctrl.todoService.CreateTodo(todoItem.Title, todoItem.Content, todoItem.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusCreated, response.TodoResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{}})
}

func (ctrl *controller) GetTodoList(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, response.TodoResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": "not found username"}})
		return
	}

	todoList, err := ctrl.todoService.GetTodoList(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, response.TodoResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": todoList}})
}

func (ctrl *controller) DeleteTodoItem(c *gin.Context) {
	todoId := c.Param("todoId")
	objId, _ := primitive.ObjectIDFromHex(todoId)

	err := ctrl.todoService.DeleteTodoItem(objId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, response.TodoResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}})
}

func (ctrl *controller) EditTodoItem(c *gin.Context) {
	todoId := c.Param("todoId")
	objId, _ := primitive.ObjectIDFromHex(todoId)

	var todoItem todoModel.Todo
	if err := c.BindJSON(&todoItem); err != nil {
		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	err := ctrl.todoService.EditTodoItem(objId, todoItem.Content, todoItem.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, response.TodoResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{}})
}
