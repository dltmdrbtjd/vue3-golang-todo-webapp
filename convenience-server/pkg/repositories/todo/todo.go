package todo

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repostitory) CreateTodo(ctx context.Context, title string, status string, content string, username string) error {

	t := time.Now()
	timeFormat := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())

	res, err := r.todoCollection.InsertOne(ctx,
		bson.M{
			"title":     title,
			"status":    status,
			"content":   content,
			"username":  username,
			"create_at": timeFormat,
		})

	if err != nil {
		return err
	}
	logrus.Info("new todo item create. id:", res.InsertedID)
	return nil
}
