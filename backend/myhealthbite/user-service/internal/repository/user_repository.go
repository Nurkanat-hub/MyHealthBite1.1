package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`

	Goal    string  `bson:"goal,omitempty"`
	Height  float32 `bson:"height,omitempty"`
	Weight  float32 `bson:"weight,omitempty"`
	Address string  `bson:"address,omitempty"`
	Phone   string  `bson:"phone,omitempty"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(col *mongo.Collection) *UserRepository {
	return &UserRepository{Collection: col}
}

func (r *UserRepository) CreateUser(ctx context.Context, user User) error {
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}
func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	count, err := r.Collection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (r *UserRepository) FindByID(ctx context.Context, id primitive.ObjectID) (User, error) {
	var user User
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}
func (r *UserRepository) UpdateUserProfile(ctx context.Context, id primitive.ObjectID, updates User) (User, error) {
	update := bson.M{
		"$set": bson.M{
			"name":    updates.Name,
			"goal":    updates.Goal,
			"height":  updates.Height,
			"weight":  updates.Weight,
			"address": updates.Address,
			"phone":   updates.Phone,
		},
	}

	_, err := r.Collection.UpdateByID(ctx, id, update)
	if err != nil {
		return User{}, err
	}

	return r.FindByID(ctx, id)
}
func (r *UserRepository) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
