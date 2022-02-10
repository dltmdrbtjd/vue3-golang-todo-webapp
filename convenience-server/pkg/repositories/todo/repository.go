package todo

import (
	"context"

	"github.com/Convenience-Tools/convenience-server/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	todoCollectionName = "todo"
)

type Repository interface {
	CreateTodo(ctx context.Context, title string, status string, content string, username string) error
}

type repostitory struct {
	todoCollection *mongo.Collection
}

func NewRepository(todoCollection *mongo.Collection) Repository {
	return &repostitory{
		todoCollection: todoCollection,
	}
}

func CreateRepository() Repository {
	return NewRepository(
		db.GetDB().Collection(todoCollectionName),
	)
}
