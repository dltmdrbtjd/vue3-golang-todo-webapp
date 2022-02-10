package db

import (
	"context"
	"fmt"

	"github.com/Convenience-Tools/convenience-server/pkg/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

func GetDB() *mongo.Database {
	return database
}

func Initialize(ctx context.Context) (err error) {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", config.DBUserName, config.DBPassword, config.DBPort)
	fmt.Printf("Trying to connect %s...", uri)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://dltmdrbtjd:ss29347718@cluster0.5w9nj.mongodb.net/golangDB?retryWrites=true&w=majority"))
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	database = client.Database(config.DBName)
	return nil
}

func Close(ctx context.Context) error {
	return client.Disconnect(ctx)
}
