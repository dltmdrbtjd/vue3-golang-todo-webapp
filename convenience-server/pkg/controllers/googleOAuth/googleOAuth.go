package googleOAuth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Convenience-Tools/convenience-server/pkg/models/user"
	"github.com/Convenience-Tools/convenience-server/pkg/services/googleOAuth"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type controller struct {
	googleOAuthService googleOAuth.Service
}

type Controller interface {
	GoogleLogin(c *gin.Context)
	GoogleLoginCallback(c *gin.Context)
	GetGoogleUserInfo(c *gin.Context)
	GoogleTokenVerification(c *gin.Context)
}

func NewController(googleOAuthService googleOAuth.Service) *controller {
	return &controller{
		googleOAuthService: googleOAuthService,
	}
}

func CreateController() *controller {
	return NewController(googleOAuth.CreateService())
}

func (ctrl *controller) GoogleLogin(c *gin.Context) {
	url := ctrl.googleOAuthService.GoogleLogin()
	c.JSON(http.StatusOK, map[string]string{"data": url})
}

func (ctrl *controller) GoogleLoginCallback(c *gin.Context) {
	token, err := ctrl.googleOAuthService.GoogleLoginCallback(c.Request.FormValue("code"))
	if err != nil {
		fmt.Println("could not get token \n", err.Error())
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("could not create reuqest \n", err.Error())
		c.JSON(http.StatusNonAuthoritativeInfo, map[string]string{"error": "unexcepted access token!"})
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read response body \n", err.Error())
		return
	}

	var userinfo user.GoogleUserInfo
	userinfo.Token = token
	err = json.Unmarshal(content, &userinfo)
	if err != nil {
		log.Println(err)
	}

	err = ctrl.googleOAuthService.SaveGoogleUserInfo(&userinfo)
	if err != nil {
		logrus.Errorln(err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"data": token.AccessToken})
}

func (ctrl *controller) GoogleTokenVerification(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		logrus.Errorln("not found user token controller")
		c.JSON(http.StatusNotFound, map[string]string{"error": "not found user email"})
	}

	_, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		fmt.Println("could not create reuqest \n", err.Error())
		c.JSON(http.StatusNonAuthoritativeInfo, map[string]string{"error": "unexcepted access token!"})
		return
	}

	c.JSON(http.StatusOK, map[string]bool{"verification": true})
}

func (ctrl *controller) GetGoogleUserInfo(c *gin.Context) {
	userEmail := c.Param("useremail")
	if userEmail == "" {
		logrus.Errorln("not found user email controller")
		c.JSON(http.StatusNotFound, map[string]string{"error": "not found user email"})
	}

	userInfo, err := ctrl.googleOAuthService.GetGoogleUserInfo(userEmail)
	if err != nil {
		logrus.Errorln("not found user info controller", err.Error())
		c.JSON(http.StatusNotFound, map[string]string{"error": "user info not found"})
	}

	c.JSON(http.StatusOK, map[string]interface{}{"data": userInfo})
}
