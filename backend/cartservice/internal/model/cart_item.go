package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   string             `bson:"user_id"`
	DishID   string             `bson:"dish_id"`
	Quantity int32              `bson:"quantity"`
	AddedAt  time.Time          `bson:"added_at"`
}
