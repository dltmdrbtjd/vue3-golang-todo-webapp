package controllers

import (
	"net/http"

	"github.com/Convenience-Tools/convenience-server/pkg/models"
	"github.com/Convenience-Tools/convenience-server/pkg/response"
	"github.com/Convenience-Tools/convenience-server/pkg/services"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&todo); validationErr != nil {
		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	newTodo, err := services.CreateTodoService(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, response.TodoResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": newTodo}})
}

func GetTodoList(c *gin.Context) {
	username := c.Param("userName")

	todoList, err := services.GetTodoListService(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusCreated, response.TodoResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": todoList}})
}

func DeleteTodo(c *gin.Context) {
	var userName string

	if err := c.BindJSON(&userName); err != nil {
		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&userName); validationErr != nil {
		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	err := services.DeleteTodoService(userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, response.TodoResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{}})
}
