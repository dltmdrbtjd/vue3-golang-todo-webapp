package user

import (
	"context"

	"github.com/Convenience-Tools/convenience-server/pkg/models/user"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) SaveGoogleUserInfo(ctx context.Context, userInfo *user.GoogleUserInfo) error {
	_, err := r.userCollection.InsertOne(ctx, userInfo)
	if err != nil {
		logrus.Errorln(err.Error())
		return err
	}

	return err
}

func (r *repository) GetGoogleUserInfo(ctx context.Context, token string) (user.GoogleUserInfo, error) {
	var userInfo user.GoogleUserInfo
	err := r.userCollection.FindOne(ctx, bson.M{"accesstoken": token}).Decode(&userInfo)
	if err != nil {
		logrus.Errorln("not found user info:", err.Error())
	}

	return userInfo, err
}
