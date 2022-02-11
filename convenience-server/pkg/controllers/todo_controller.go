package controllers

// import (
// 	"net/http"

// 	"github.com/Convenience-Tools/convenience-server/pkg/models"
// 	"github.com/Convenience-Tools/convenience-server/pkg/response"
// 	"github.com/Convenience-Tools/convenience-server/pkg/services"
// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func EditTodo(c *gin.Context) {
// 	todoId := c.Param("objId")
// 	objId, _ := primitive.ObjectIDFromHex(todoId)
// 	var editTodo models.UpdateTodo

// 	if err := c.BindJSON(&editTodo); err != nil {
// 		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 		return
// 	}

// 	updatedTodo, err := services.EditTodoService(objId, editTodo)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.TodoResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedTodo}})
// }

// func TodoStatusChange(c *gin.Context) {
// 	todoId := c.Param("objId")
// 	objId, _ := primitive.ObjectIDFromHex(todoId)
// 	var currentStatus models.TodoCurrentStatus

// 	if err := c.BindJSON(&currentStatus); err != nil {
// 		c.JSON(http.StatusBadRequest, response.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 		return
// 	}

// 	status, err := services.TodoStatusChangeService(objId, currentStatus)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.TodoResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": status}})
// }
