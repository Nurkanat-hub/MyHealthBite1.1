package handler

import (
	"context"
	"time"

	"MyHealthBite/stats-service/internal/repository"
	"MyHealthBite/stats-service/proto"
)

type StatsServer struct {
	proto.UnimplementedStatsServiceServer
	Repo *repository.StatsRepository
}

func NewStatsServer(repo *repository.StatsRepository) *StatsServer {
	return &StatsServer{Repo: repo}
}

func (s *StatsServer) InitStats(ctx context.Context, req *proto.InitStatsRequest) (*proto.StatsResponse, error) {
	stats := repository.Stats{
		UserID:          req.UserId,
		TargetCalories:  req.TargetCalories,
		TargetWaterML:   req.TargetWaterMl,
		CurrentCalories: 0,
		CurrentWaterML:  0,
		UpdatedAt:       time.Now(),
	}
	err := s.Repo.InitStats(ctx, stats)
	if err != nil {
		return nil, err
	}
	return toProto(stats), nil
}

func (s *StatsServer) GetStats(ctx context.Context, req *proto.UserIdRequest) (*proto.StatsResponse, error) {
	stats, err := s.Repo.GetStatsByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return toProto(stats), nil
}

func (s *StatsServer) UpdateStats(ctx context.Context, req *proto.UpdateStatsRequest) (*proto.StatsResponse, error) {
	stats, err := s.Repo.UpdateStats(ctx, req.UserId, req.DeltaCalories, req.DeltaWaterMl)
	if err != nil {
		return nil, err
	}
	return toProto(stats), nil
}

func (s *StatsServer) ResetDailyStats(ctx context.Context, req *proto.UserIdRequest) (*proto.StatsResponse, error) {
	stats, err := s.Repo.ResetDailyStats(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return toProto(stats), nil
}

func (s *StatsServer) DeleteStatsByUserId(ctx context.Context, req *proto.UserIdRequest) (*proto.Empty, error) {
	err := s.Repo.DeleteByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

// Вспомогательная функция для преобразования модели в proto
func toProto(s repository.Stats) *proto.StatsResponse {
	return &proto.StatsResponse{
		UserId:          s.UserID,
		TargetCalories:  s.TargetCalories,
		CurrentCalories: s.CurrentCalories,
		TargetWaterMl:   s.TargetWaterML,
		CurrentWaterMl:  s.CurrentWaterML,
		UpdatedAt:       s.UpdatedAt.Format(time.RFC3339),
	}
}
