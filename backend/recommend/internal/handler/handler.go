package handler

import (
	"context"

	pb "recommend/internal/proto"
)

type Handler struct {
	pb.UnimplementedRecommendationServiceServer
	svc RecommendationLogic
}

type RecommendationLogic interface {
	GetPopularDishes() ([]*pb.Dish, error)
}

func NewHandler(svc RecommendationLogic) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetPopularDishes(ctx context.Context, _ *pb.PopularRequest) (*pb.PopularResponse, error) {
	dishes, err := h.svc.GetPopularDishes()
	if err != nil {
		return nil, err
	}
	return &pb.PopularResponse{Dishes: dishes}, nil
}
