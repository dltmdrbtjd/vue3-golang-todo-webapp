package user

import (
	"context"

	"github.com/Convenience-Tools/convenience-server/pkg/db"
	"github.com/Convenience-Tools/convenience-server/pkg/models/user"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollectionName = "user"
)

type Repository interface {
	SaveGoogleUserInfo(ctx context.Context, userInfo *user.GoogleUserInfo) error
	GetGoogleUserInfo(ctx context.Context, userEmail string) (user.UserInfo, error)
}

type repository struct {
	userCollection *mongo.Collection
}

func NewRepository(userCollection *mongo.Collection) Repository {
	return &repository{
		userCollection: userCollection,
	}
}

func CreateRepository() Repository {
	return NewRepository(
		db.GetDB().Collection(userCollectionName),
	)
}
