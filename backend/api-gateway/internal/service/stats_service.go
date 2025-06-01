package service

import (
	"context"
	"log"
	"time"

	statsProto "MyHealthBite/stats-service/proto"

	"google.golang.org/grpc"
)

var StatsClient statsProto.StatsServiceClient

func InitStatsServiceClient() {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("❌ Could not connect to stats-service: %v", err)
	}

	StatsClient = statsProto.NewStatsServiceClient(conn)
	log.Println("✅ Connected to stats-service via gRPC")
}

// Новый конструктор клиента по соединению
func NewStatsClient(conn *grpc.ClientConn) statsProto.StatsServiceClient {
	return statsProto.NewStatsServiceClient(conn)
}

// Вспомогательная функция для создания запроса GetStats
func NewUserIdRequest(userId string) *statsProto.UserIdRequest {
	return &statsProto.UserIdRequest{UserId: userId}
}

// Пример вызова GetStats через глобальный клиент
func GetStats(ctx context.Context, userId string) (*statsProto.StatsResponse, error) {
	return StatsClient.GetStats(ctx, NewUserIdRequest(userId))
}
