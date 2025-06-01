package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dish struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name"`
	Description     string             `bson:"description"`
	Category        string             `bson:"category"`
	Calories        int32              `bson:"calories"`
	Proteins        int32              `bson:"proteins"`
	Fats            int32              `bson:"fats"`
	Carbs           int32              `bson:"carbs"`
	Ingredients     []string           `bson:"ingredients"`
	CookTimeMinutes int32              `bson:"cook_time_minutes"`
	Price           int32              `bson:"price"`
}
