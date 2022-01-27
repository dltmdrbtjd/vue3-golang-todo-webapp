package routers

import (
	"github.com/Convenience-Tools/convenience-server/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func TestRoutes(route *gin.Engine) {
	route.GET("/", controllers.TestController)
}
