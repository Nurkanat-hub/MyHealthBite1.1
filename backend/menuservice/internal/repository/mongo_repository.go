package repository

import (
	"context"
	"menuservice/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dishCollection *mongo.Collection

func SetCollection(c *mongo.Collection) {
	dishCollection = c
}

func InitMongo(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	dishCollection = client.Database("myhealthbite").Collection("menu")
	return nil
}

type DishRepository interface {
	Create(d *model.Dish) (*model.Dish, error)
	GetByID(id string) (*model.Dish, error)
	Update(d *model.Dish) (*model.Dish, error)
	Delete(id string) error
	List() ([]*model.Dish, error)
}

type mongoRepo struct{}

func NewMongoRepository() DishRepository {
	return &mongoRepo{}
}

func (r *mongoRepo) Create(d *model.Dish) (*model.Dish, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := dishCollection.InsertOne(ctx, d)
	if err != nil {
		return nil, err
	}
	d.ID = res.InsertedID.(primitive.ObjectID)
	return d, nil
}

func (r *mongoRepo) GetByID(id string) (*model.Dish, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var d model.Dish
	err = dishCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *mongoRepo) Update(d *model.Dish) (*model.Dish, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := dishCollection.ReplaceOne(ctx, bson.M{"_id": d.ID}, d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *mongoRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = dishCollection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (r *mongoRepo) List() ([]*model.Dish, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := dishCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var dishes []*model.Dish
	for cursor.Next(ctx) {
		var d model.Dish
		if err := cursor.Decode(&d); err != nil {
			return nil, err
		}
		dishes = append(dishes, &d)
	}
	return dishes, nil
}
