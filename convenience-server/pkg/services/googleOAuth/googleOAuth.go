package googleOAuth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"os"

	"github.com/Convenience-Tools/convenience-server/pkg/models/user"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
)

func GetToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.RawStdEncoding.EncodeToString(b)
}

func (s *service) GoogleLogin() string {
	urlToken := GetToken()
	uri := googleOauthConfig.AuthCodeURL(urlToken)
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

func (s *service) SaveGoogleUserInfo(userInfo *user.GoogleUserInfo) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := s.userRepository.SaveGoogleUserInfo(ctx, userInfo)
	if err != nil {
		logrus.Errorln(err.Error())
		return err
	}
	return err
}

func (s *service) GetGoogleUserInfo(token string) (*user.GoogleUserInfo, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	userInfo, err := s.userRepository.GetGoogleUserInfo(ctx, token)
	if err != nil {
		logrus.Errorln("not found user info service:", err.Error())
		return nil, err
	}

	return &userInfo, err
}
