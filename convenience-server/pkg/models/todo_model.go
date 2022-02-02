package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Title    string             `json:"title,omitempty" validate:"required"`
	Content  string             `json:"content,omitempty" validate:"required"`
	Status   string             `json:"status,omitempty"`
	UserName string             `json:"username,omitempty" validate:"required"`
}
