package googleOAuth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Convenience-Tools/convenience-server/pkg/models/user"
	"github.com/Convenience-Tools/convenience-server/pkg/services/oauth"
	"github.com/gin-gonic/gin"
)

type controller struct {
	oauthService oauth.Service
}

type Controller interface {
	GoogleLogin(c *gin.Context)
	GoogleLoginCallback(c *gin.Context)
}

func NewController(oauthService oauth.Service) *controller {
	return &controller{
		oauthService: oauthService,
	}
}

func CreateController() *controller {
	return NewController(oauth.CreateService())
}

func (ctrl *controller) GoogleLogin(c *gin.Context) {
	url := ctrl.oauthService.GoogleLogin()
	c.JSON(http.StatusOK, map[string]string{"data": url})
}

func (ctrl *controller) GoogleLoginCallback(c *gin.Context) {
	token, err := ctrl.oauthService.GoogleLoginCallback(c.Request.FormValue("code"))
	if err != nil {
		fmt.Println("could not get token \n", err.Error())
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("could not create reuqest \n", err.Error())
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read response body \n", err.Error())
		return
	}

	var userinfo user.UserInfo
	err = json.Unmarshal(content, &userinfo)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{"email": userinfo.Email})
}
