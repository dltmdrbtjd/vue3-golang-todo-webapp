package todo

import "github.com/Convenience-Tools/convenience-server/pkg/repositories/todo"

type Service interface {
	CreateTodo(title string, status string, content string, username string) error
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
