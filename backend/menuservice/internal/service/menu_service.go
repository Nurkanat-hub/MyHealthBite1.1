package service

import (
	"context"
	"menuservice/internal/model"
	"menuservice/internal/repository"
	proto "menuservice/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuService struct {
	proto.UnimplementedMenuServiceServer
	Repo repository.DishRepository
}

func NewMenuService(repo repository.DishRepository) *MenuService {
	return &MenuService{Repo: repo}
}

func (s *MenuService) CreateDish(ctx context.Context, req *proto.CreateDishRequest) (*proto.Dish, error) {
	dish := &model.Dish{
		Name:            req.Name,
		Description:     req.Description,
		Category:        req.Category,
		Calories:        req.Calories,
		Proteins:        req.Proteins,
		Fats:            req.Fats,
		Carbs:           req.Carbs,
		Ingredients:     req.Ingredients,
		CookTimeMinutes: req.CookTimeMinutes,
		Price:           req.Price,
	}

	created, err := s.Repo.Create(dish)
	if err != nil {
		return nil, err
	}
	return toProto(created), nil
}

func (s *MenuService) GetDishById(ctx context.Context, req *proto.DishIdRequest) (*proto.Dish, error) {
	dish, err := s.Repo.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return toProto(dish), nil
}

func (s *MenuService) UpdateDish(ctx context.Context, req *proto.UpdateDishRequest) (*proto.Dish, error) {
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	dish := &model.Dish{
		ID:              oid,
		Name:            req.Name,
		Description:     req.Description,
		Category:        req.Category,
		Calories:        req.Calories,
		Proteins:        req.Proteins,
		Fats:            req.Fats,
		Carbs:           req.Carbs,
		Ingredients:     req.Ingredients,
		CookTimeMinutes: req.CookTimeMinutes,
		Price:           req.Price,
	}

	updated, err := s.Repo.Update(dish)
	if err != nil {
		return nil, err
	}
	return toProto(updated), nil
}

func (s *MenuService) DeleteDish(ctx context.Context, req *proto.DishIdRequest) (*proto.DeleteResponse, error) {
	err := s.Repo.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteResponse{Message: "Deleted successfully"}, nil
}

func (s *MenuService) GetAllDishes(ctx context.Context, _ *proto.Empty) (*proto.DishList, error) {
	items, err := s.Repo.List()
	if err != nil {
		return nil, err
	}

	var result []*proto.Dish
	for _, item := range items {
		result = append(result, toProto(item))
	}
	return &proto.DishList{Dishes: result}, nil
}

func toProto(d *model.Dish) *proto.Dish {
	return &proto.Dish{
		Id:              d.ID.Hex(),
		Name:            d.Name,
		Description:     d.Description,
		Category:        d.Category,
		Calories:        d.Calories,
		Proteins:        d.Proteins,
		Fats:            d.Fats,
		Carbs:           d.Carbs,
		Ingredients:     d.Ingredients,
		CookTimeMinutes: d.CookTimeMinutes,
		Price:           d.Price,
	}
}
