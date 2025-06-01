package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	DB *mongo.Collection
}

func ConnectMongo() (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://Nurkanat:Nurkanat05@myhealthbite.sfcpxor.mongodb.net/myhealthbite_m",
	))
	if err != nil {
		return nil, err
	}

	db := client.Database("myhealthbite_menu").Collection("dishes")
	store := &Storage{DB: db}

	// Автовставка блюд при первом запуске
	count, err := db.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if count == 0 {
		log.Println("dishes collection is empty — inserting default dishes...")
		err := store.insertDefaultPopularDishes(ctx)
		if err != nil {
			return nil, err
		}
	}

	return store, nil
}

func (s *Storage) GetPopularDishes() ([]map[string]interface{}, error) {
	ctx := context.Background()
	filter := bson.M{"popular": true}
	cursor, err := s.DB.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (s *Storage) insertDefaultPopularDishes(ctx context.Context) error {
	dishes := []interface{}{
		bson.M{"id": "101", "name": "Шашлык из курицы", "description": "Сочный шашлык", "category": "ужин", "calories": 520, "proteins": 35, "fats": 25, "carbs": 5, "ingredients": []string{"курица", "лук"}, "cook_time_minutes": 40, "price": 2200.0, "popular": true},
		bson.M{"id": "102", "name": "Овсянка с ягодами", "description": "Полезный завтрак", "category": "завтрак", "calories": 290, "proteins": 8, "fats": 6, "carbs": 45, "ingredients": []string{"овсянка", "ягоды"}, "cook_time_minutes": 10, "price": 850.0, "popular": true},
		bson.M{"id": "103", "name": "Паста болоньезе", "description": "Мясной соус и сыр", "category": "ужин", "calories": 580, "proteins": 28, "fats": 22, "carbs": 65, "ingredients": []string{"спагетти", "фарш"}, "cook_time_minutes": 25, "price": 1850.0, "popular": true},
		bson.M{"id": "104", "name": "Греческий салат", "description": "Овощи и фета", "category": "перекус", "calories": 210, "proteins": 6, "fats": 12, "carbs": 18, "ingredients": []string{"помидоры", "огурцы"}, "cook_time_minutes": 10, "price": 1200.0, "popular": true},
		bson.M{"id": "105", "name": "Плов", "description": "Сытное блюдо с мясом", "category": "обед", "calories": 740, "proteins": 28, "fats": 30, "carbs": 80, "ingredients": []string{"рис", "мясо"}, "cook_time_minutes": 50, "price": 2500.0, "popular": true},
		bson.M{"id": "106", "name": "Суп с фрикадельками", "description": "Домашний суп", "category": "обед", "calories": 340, "proteins": 18, "fats": 10, "carbs": 30, "ingredients": []string{"фарш", "картофель"}, "cook_time_minutes": 30, "price": 1500.0, "popular": true},
		bson.M{"id": "107", "name": "Курица терияки", "description": "Сладкий соус и рис", "category": "ужин", "calories": 550, "proteins": 30, "fats": 15, "carbs": 50, "ingredients": []string{"курица", "соус"}, "cook_time_minutes": 30, "price": 1950.0, "popular": true},
		bson.M{"id": "108", "name": "Смузи банан-клубника", "description": "Фруктовый напиток", "category": "перекус", "calories": 160, "proteins": 3, "fats": 2, "carbs": 35, "ingredients": []string{"банан", "клубника"}, "cook_time_minutes": 5, "price": 700.0, "popular": true},
		bson.M{"id": "109", "name": "Том Ям", "description": "Острый суп с креветками", "category": "обед", "calories": 420, "proteins": 22, "fats": 15, "carbs": 25, "ingredients": []string{"креветки", "кокос"}, "cook_time_minutes": 20, "price": 2300.0, "popular": true},
		bson.M{"id": "110", "name": "Бургер", "description": "Котлета, сыр, соус", "category": "обед", "calories": 650, "proteins": 30, "fats": 35, "carbs": 40, "ingredients": []string{"булка", "мясо"}, "cook_time_minutes": 15, "price": 2000.0, "popular": true},
	}

	_, err := s.DB.InsertMany(ctx, dishes)
	return err
}
