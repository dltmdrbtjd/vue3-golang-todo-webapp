package user

import (
	"context"

	"github.com/Convenience-Tools/convenience-server/pkg/models/user"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/oauth2"
)

func (r *repository) SaveGoogleUserInfo(ctx context.Context, userInfo *user.GoogleUserInfo) error {
	_, err := r.userCollection.InsertOne(ctx, userInfo)
	if err != nil {
		logrus.Errorln(err.Error())
		return err
	}

	return err
}

func (r *repository) GetGoogleUserInfo(ctx context.Context, userEmail string) (user.UserInfo, error) {
	var userInfo user.UserInfo
	err := r.userCollection.FindOne(ctx, bson.M{"email": userEmail}).Decode(&userInfo)
	if err != nil {
		logrus.Errorln("not found user info:", err.Error())
	}

	return userInfo, err
}

func (r *repository) GoogleTokenVerification(ctx context.Context, userEmail string) (*oauth2.Token, error) {
	var userInfo user.GoogleUserInfo
	err := r.userCollection.FindOne(ctx, bson.M{"email": userEmail}).Decode(&userInfo)
	if err != nil {
		logrus.Errorln("not found user info:", err.Error())
	}

	return userInfo.Token, err
}
