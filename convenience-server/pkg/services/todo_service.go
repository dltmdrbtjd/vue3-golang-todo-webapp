package services

import (
	"context"
	"encoding/json"
	"fmt"
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

	t := time.Now()
	timeFormat := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())

	newTodo := models.Todo{
		Id:       primitive.NewObjectID(),
		Title:    todo.Title,
		Status:   "false",
		Content:  todo.Content,
		UserName: todo.UserName,
		CreateAt: timeFormat,
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

func DeleteTodoService(objId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := todoCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		log.Fatal(err)
	}

	if result.DeletedCount < 1 {
		log.Fatal("User todo specified Id not found")
	}

	return err
}

func EditTodoService(objId primitive.ObjectID, editTodo models.UpdateTodo) (*models.UpdateTodo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	currentTodo := findOneTodoItem(objId)

	if editTodo.Content == "" {
		editTodo.Content = currentTodo.Content
	}
	if editTodo.Title == "" {
		editTodo.Title = currentTodo.Title
	}
	if editTodo.UserName == "" {
		editTodo.UserName = currentTodo.UserName
	}

	update := bson.M{"title": editTodo.Title, "content": editTodo.Content, "username": editTodo.UserName}
	result, err := todoCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		log.Fatal(err)
	}

	var updatedTodo *models.UpdateTodo
	if result.MatchedCount == 1 {
		err := todoCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedTodo)
		if err != nil {
			log.Fatal(err)
		}
	}

	return updatedTodo, err
}

func findOneTodoItem(objId primitive.ObjectID) *models.Todo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var todo *models.Todo
	err := todoCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}

	return todo
}

func TodoStatusChangeService(objId primitive.ObjectID, status models.TodoCurrentStatus) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// var newStatus models.TodoStatus
	json.Marshal(status.Status)

	updateStatus := bson.M{"status": string(status.Status)}

	_, err := todoCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updateStatus})
	if err != nil {
		log.Fatal(err)
	}

	return status.Status, err
}
