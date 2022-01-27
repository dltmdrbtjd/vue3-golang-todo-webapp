package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestController(c *gin.Context) {
	test := "test hello world"

	c.JSON(http.StatusOK, gin.H{"data": test})
}
