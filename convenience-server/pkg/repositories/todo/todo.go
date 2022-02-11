package todo

import (
	"context"
	"log"
	"time"

	"github.com/Convenience-Tools/convenience-server/pkg/models/todo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repostitory) CreateTodo(ctx context.Context, title string, content string, username string) error {

	t := time.Now()

	res, err := r.todoCollection.InsertOne(ctx,
		bson.M{
			"title":     title,
			"status":    "ready",
			"content":   content,
			"username":  username,
			"create_at": t,
		})

	if err != nil {
		return err
	}
	logrus.Info("new todo item create. id:", res.InsertedID)
	return nil
}

func (r *repostitory) GetTodoList(ctx context.Context, username string) ([]todo.Todo, error) {
	findOptions := options.FindOptions{}

	cur, err := r.todoCollection.Find(ctx,
		bson.M{
			"username": username,
		},
		&findOptions,
	)
	if err != nil {
		logrus.Fatalln("get todolist fatal error :", err.Error())
	}

	var todoList []todo.Todo
	if err = cur.All(ctx, &todoList); err != nil {
		log.Fatalln("get todolist failed to decode todolist:", err.Error())
		return nil, err
	}

	return todoList, nil
}

func (r *repostitory) DeleteTodoItem(ctx context.Context, todoId primitive.ObjectID) error {
	res, err := r.todoCollection.DeleteOne(ctx,
		bson.M{
			"_id": todoId,
		},
	)

	if err != nil {
		logrus.Fatalln("delete todo item fatal error :", err.Error())
	}

	if res.DeletedCount < 1 {
		log.Fatalln("user todo specified id not found:", err.Error())
	}

	return err
}

func (r *repostitory) EditTodoItem(ctx context.Context, todoId primitive.ObjectID, content string, title string) error {
	todoItem := r.findOneTodoItem(ctx, todoId)

	if content == "" {
		content = todoItem.Content
	}

	if title == "" {
		title = todoItem.Title
	}

	updateTodoItem := bson.M{"content": content, "title": title}
	_, err := r.todoCollection.UpdateOne(ctx, bson.M{"_id": todoId}, updateTodoItem)
	if err != nil {
		logrus.Errorln("todo item failed update:", err.Error())
		return err
	}

	return err
}

func (r *repostitory) findOneTodoItem(ctx context.Context, todoId primitive.ObjectID) todo.Todo {
	var todoItem todo.Todo
	err := r.todoCollection.FindOne(ctx, bson.M{"_id": todoId}).Decode(&todoItem)
	if err != nil {
		logrus.Errorln("not found one todo item:", err.Error())
	}

	return todoItem
}
