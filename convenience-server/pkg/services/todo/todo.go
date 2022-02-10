package todo

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (s *service) CreateTodo(title string, status string, content string, username string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := s.todoRepository.CreateTodo(ctx, title, status, content, username)
	if err != nil {
		logrus.Errorln("CreateTodo error: ", err.Error())
		return err
	}

	return nil
}
