package todo

import (
	"context"

	"github.com/Convenience-Tools/convenience-server/pkg/models/todo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) CreateTodo(title string, content string, username string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := s.todoRepository.CreateTodo(ctx, title, content, username)
	if err != nil {
		logrus.Errorln("create todo error: ", err.Error())
		return err
	}

	return nil
}

func (s *service) GetTodoList(username string) ([]todo.Todo, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	todoList, err := s.todoRepository.GetTodoList(ctx, username)
	if err != nil {
		logrus.Errorln("get todolist error: ", err.Error())
		return nil, err
	}

	return todoList, err
}

func (s *service) DeleteTodoItem(todoId primitive.ObjectID) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := s.todoRepository.DeleteTodoItem(ctx, todoId)
	if err != nil {
		logrus.Errorln("delete todo item error:", err.Error())
	}

	return err
}

func (s *service) EditTodoItem(todoId primitive.ObjectID, content string, title string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := s.todoRepository.EditTodoItem(ctx, todoId, content, title)
	if err != nil {
		logrus.Errorln("edit todo item error:", err.Error())
	}

	return err
}
