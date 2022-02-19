package oauth

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = oauth2.Config{
		RedirectURL:  "http://localhost:8081/google-login/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	RandomState = "random"
)

func (s *service) GoogleLogin() string {
	uri := googleOauthConfig.AuthCodeURL(RandomState)
	if uri == "" {
		logrus.Errorln("not found google oauth uri!")
		return ""
	}

	return uri
}

func (s *service) GoogleLoginCallback(code string) (*oauth2.Token, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	token, err := googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return token, err
}
