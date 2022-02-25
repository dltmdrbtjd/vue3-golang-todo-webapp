package base

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	CreateAt time.Time          `json:"create_at" bson:"create_at"`
	UpdateAt time.Time          `json:"update_at" bson:"update_at"`
}
