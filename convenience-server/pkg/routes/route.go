package routes

import (
	"net/http"

	"github.com/Convenience-Tools/convenience-server/pkg/routes/googleLogin"
	"github.com/Convenience-Tools/convenience-server/pkg/routes/todo"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"response": "pong",
		})
	})
	rootGroup := r.Group("/")
	todo.Register(rootGroup, nil)
	googleLogin.Register(rootGroup, nil)
}
