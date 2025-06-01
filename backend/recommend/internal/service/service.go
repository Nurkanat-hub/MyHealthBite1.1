package service

import (
	pb "recommend/internal/proto"
	"recommend/internal/storage"
)

type RecommendationService struct {
	store *storage.Storage
}

func NewRecommendationService(store *storage.Storage) *RecommendationService {
	return &RecommendationService{store: store}
}

func (s *RecommendationService) GetPopularDishes() ([]*pb.Dish, error) {
	data, err := s.store.GetPopularDishes()
	if err != nil {
		return nil, err
	}

	var dishes []*pb.Dish
	for _, item := range data {
		dishes = append(dishes, &pb.Dish{
			Id:              item["id"].(string),
			Name:            item["name"].(string),
			Description:     item["description"].(string),
			Category:        item["category"].(string),
			Calories:        int32(item["calories"].(int32)),
			Proteins:        int32(item["proteins"].(int32)),
			Fats:            int32(item["fats"].(int32)),
			Carbs:           int32(item["carbs"].(int32)),
			Ingredients:     toStringSlice(item["ingredients"]),
			CookTimeMinutes: int32(item["cook_time_minutes"].(int32)),
			Price:           float32(item["price"].(float64)),
		})
	}
	return dishes, nil
}

func toStringSlice(val interface{}) []string {
	raw := val.([]interface{})
	result := make([]string, len(raw))
	for i, v := range raw {
		result[i] = v.(string)
	}
	return result
}
