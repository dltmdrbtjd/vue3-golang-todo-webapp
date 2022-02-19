package oauth

import (
	"golang.org/x/oauth2"
)

type Service interface {
	GoogleLogin() string
	GoogleLoginCallback(code string) (*oauth2.Token, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func CreateService() Service {
	return NewService()
}
