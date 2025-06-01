package service

import (
	"context"
	"time"

	"cartservice/internal/model"
	"cartservice/internal/repository"
	pb "cartservice/proto"
)

type CartService struct {
	pb.UnimplementedCartServiceServer
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.Empty, error) {
	item := &model.CartItem{
		UserID:   req.UserId,
		DishID:   req.DishId,
		Quantity: req.Quantity,
		AddedAt:  time.Now(),
	}
	if err := s.repo.AddItem(item); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *CartService) RemoveFromCart(ctx context.Context, req *pb.RemoveFromCartRequest) (*pb.Empty, error) {
	if err := s.repo.RemoveItem(req.UserId, req.DishId); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.CartList, error) {
	items, err := s.repo.GetCart(req.UserId)
	if err != nil {
		return nil, err
	}

	var result []*pb.CartItem
	for _, item := range items {
		result = append(result, &pb.CartItem{
			Id:       item.ID.Hex(),
			UserId:   item.UserID,
			DishId:   item.DishID,
			Quantity: item.Quantity,
			AddedAt:  item.AddedAt.Format(time.RFC3339),
		})
	}

	return &pb.CartList{Items: result}, nil
}
