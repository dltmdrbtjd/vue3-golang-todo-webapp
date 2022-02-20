package googleOAuth

import (
	"github.com/Convenience-Tools/convenience-server/pkg/models/user"
	userRepo "github.com/Convenience-Tools/convenience-server/pkg/repositories/user"
	"golang.org/x/oauth2"
)

type Service interface {
	GoogleLogin() string
	GoogleLoginCallback(code string) (*oauth2.Token, error)
	SaveGoogleUserInfo(userInfo *user.GoogleUserInfo) error
	GetGoogleUserInfo(userEmail string) (*user.UserInfo, error)
	GoogleTokenVerification(userEmail string) (*oauth2.Token, error)
}

type service struct {
	userRepository userRepo.Repository
}

func NewService(userRepository userRepo.Repository) *service {
	return &service{
		userRepository: userRepository,
	}
}

func CreateService() Service {
	return NewService(
		userRepo.CreateRepository(),
	)
}
