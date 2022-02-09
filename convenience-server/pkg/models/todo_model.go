package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Title    string             `json:"title,omitempty" binding:"required"`
	Content  string             `json:"content,omitempty" binding:"required"`
	Status   string             `json:"status,omitempty"`
	UserName string             `json:"username,omitempty" binding:"required"`
	CreateAt string             `json:"create_at,omitempty"`
}

type UpdateTodo struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserName string `json:"username"`
}

type TodoCurrentStatus struct {
	Status string `json:"status,omitempty"`
}
