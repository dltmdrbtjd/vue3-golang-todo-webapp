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

func (r *repository) CreateTodo(ctx context.Context, title string, content string, username string) (*todo.Todo, error) {

	t := time.Now()

	newTodo := bson.M{
		"title":     title,
		"status":    "ready",
		"content":   content,
		"username":  username,
		"create_at": t,
	}

	res, err := r.todoCollection.InsertOne(ctx, newTodo)

	if err != nil {
		return nil, err
	}
	logrus.Info("new todo item create. id:", res.InsertedID.(primitive.ObjectID))

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		todoItem := r.findOneTodoItem(ctx, oid)
		return &todoItem, err
	}
	return nil, err
}

func (r *repository) GetTodoList(ctx context.Context, username string) ([]todo.Todo, error) {
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

func (r *repository) DeleteTodoItem(ctx context.Context, todoId primitive.ObjectID) error {
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

func (r *repository) EditTodoItem(ctx context.Context, todoId primitive.ObjectID, content string, title string) error {
	todoItem := r.findOneTodoItem(ctx, todoId)

	if content == "" {
		content = todoItem.Content
	}

	if title == "" {
		title = todoItem.Title
	}

	updateTodoItem := bson.M{"content": content, "title": title}
	_, err := r.todoCollection.UpdateOne(ctx, bson.M{"_id": todoId}, bson.M{"$set": updateTodoItem})
	if err != nil {
		logrus.Errorln("todo item failed update:", err.Error())
		return err
	}

	return err
}

func (r *repository) findOneTodoItem(ctx context.Context, todoId primitive.ObjectID) todo.Todo {
	var todoItem todo.Todo
	err := r.todoCollection.FindOne(ctx, bson.M{"_id": todoId}).Decode(&todoItem)
	if err != nil {
		logrus.Errorln("not found one todo item:", err.Error())
	}

	return todoItem
}

func (r *repository) TodoStatusChange(ctx context.Context, todoId primitive.ObjectID, status string) error {
	updateTodoStatus := bson.M{"status": status}
	_, err := r.todoCollection.UpdateOne(ctx, bson.M{"_id": todoId}, bson.M{"$set": updateTodoStatus})
	if err != nil {
		logrus.Errorln("todo status failed change: ", err.Error())
		return err
	}

	return err
}
