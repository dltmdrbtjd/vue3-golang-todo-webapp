package services

import (
	"context"
	"log"
	"time"

	"github.com/Convenience-Tools/convenience-server/configs"
	"github.com/Convenience-Tools/convenience-server/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var todoCollection *mongo.Collection = configs.GetCollection(configs.DB, "todo")

func CreateTodoService(todo models.Todo) (models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newTodo := models.Todo{
		Id:       primitive.NewObjectID(),
		Title:    todo.Title,
		Status:   "Ready",
		Content:  todo.Content,
		UserName: todo.UserName,
	}

	_, err := todoCollection.InsertOne(ctx, newTodo)
	if err != nil {
		log.Fatal(err)
	}

	return newTodo, err
}

func GetTodoListService(userName string) ([]*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	var todoList []*models.Todo

	cur, err := todoCollection.Find(ctx, bson.M{"username": userName}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var elem models.Todo
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		todoList = append(todoList, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)
	return todoList, err
}

func DeleteTodoService(objId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := todoCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		log.Fatal(err)
	}

	if result.DeletedCount < 1 {
		log.Fatal("User todo specified Id not found")
	}

	return err
}
