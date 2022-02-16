package todo

import (
	"context"

	"github.com/Convenience-Tools/convenience-server/pkg/db"
	"github.com/Convenience-Tools/convenience-server/pkg/models/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	todoCollectionName = "todo"
)

type Repository interface {
	CreateTodo(ctx context.Context, title string, content string, username string) (*todo.Todo, error)
	GetTodoList(ctx context.Context, username string) ([]todo.Todo, error)
	DeleteTodoItem(ctx context.Context, todoId primitive.ObjectID) error
	EditTodoItem(ctx context.Context, todoId primitive.ObjectID, content string, title string) error
	TodoStatusChange(ctx context.Context, todoId primitive.ObjectID, status string) error
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
