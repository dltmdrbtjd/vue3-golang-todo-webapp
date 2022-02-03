package controllers

import (
	"net/http"

	"github.com/Convenience-Tools/convenience-server/pkg/models"
	"github.com/Convenience-Tools/convenience-server/pkg/response"
	"github.com/Convenience-Tools/convenience-server/pkg/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
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
	todoId := c.Param("objId")
	objId, _ := primitive.ObjectIDFromHex(todoId)

	err := services.DeleteTodoService(objId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, response.TodoResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}})
}

func EditTodo(c *gin.Context) {
	todoId := c.Param("objId")
	objId, _ := primitive.ObjectIDFromHex(todoId)
	var editTodo models.UpdateTodo

	if err := c.BindJSON(&editTodo); err != nil {
		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	updatedTodo, err := services.EditTodoService(objId, editTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, response.TodoResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedTodo}})
}

func TodoStatusChange(c *gin.Context) {
	todoId := c.Param("objId")
	objId, _ := primitive.ObjectIDFromHex(todoId)
	var currentStatus models.TodoStatus

	if err := c.BindJSON(&currentStatus); err != nil {
		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	status, err := services.TodoStatusChangeService(objId, currentStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, response.TodoResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": status}})
}
