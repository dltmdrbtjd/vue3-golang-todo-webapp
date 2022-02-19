package googleLogin

import (
	"github.com/Convenience-Tools/convenience-server/pkg/controllers/googleOAuth"
	"github.com/gin-gonic/gin"
)

func Register(parentGroup *gin.RouterGroup, controller googleOAuth.Controller) {
	if controller == nil {
		controller = googleOAuth.CreateController()
	}
	googleOAuthGroup := parentGroup.Group("/")
	googleOAuthGroup.Use()
	{
		googleOAuthGroup.GET("/google-login", controller.GoogleLogin)
		googleOAuthGroup.GET("/google-login/callback", controller.GoogleLoginCallback)
	}
}
