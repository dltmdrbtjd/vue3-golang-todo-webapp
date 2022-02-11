package todo

import (
	todoModel "github.com/Convenience-Tools/convenience-server/pkg/models/todo"
	"github.com/Convenience-Tools/convenience-server/pkg/repositories/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateTodo(title string, content string, username string) error
	GetTodoList(username string) ([]todoModel.Todo, error)
	DeleteTodoItem(todoId primitive.ObjectID) error
	EditTodoItem(todoId primitive.ObjectID, content string, title string) error
}

type service struct {
	todoRepository todo.Repository
}

func NewService(todoRepository todo.Repository) Service {
	return &service{
		todoRepository: todoRepository,
	}
}

func CreateService() Service {
	return NewService(
		todo.CreateRepository(),
	)
}
