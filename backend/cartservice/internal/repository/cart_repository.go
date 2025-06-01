package repository

import (
	"context"
	"time"

	"cartservice/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartRepository interface {
	AddItem(item *model.CartItem) error
	RemoveItem(userID, dishID string) error
	GetCart(userID string) ([]*model.CartItem, error)
}

type mongoCartRepo struct {
	collection *mongo.Collection
}

func NewCartRepository(col *mongo.Collection) CartRepository {
	return &mongoCartRepo{collection: col}
}

func (r *mongoCartRepo) AddItem(item *model.CartItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": item.UserID, "dish_id": item.DishID}
	update := bson.M{
		"$setOnInsert": bson.M{"added_at": time.Now()},
		"$inc":         bson.M{"quantity": item.Quantity},
	}
	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *mongoCartRepo) RemoveItem(userID, dishID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"user_id": userID, "dish_id": dishID})
	return err
}

func (r *mongoCartRepo) GetCart(userID string) ([]*model.CartItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []*model.CartItem
	for cursor.Next(ctx) {
		var item model.CartItem
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}
