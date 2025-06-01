package main

import (
	"log"
	"net"

	"recommend/internal/handler"
	"recommend/internal/service"
	"recommend/internal/storage"
	pb "recommend/internal/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := storage.ConnectMongo()
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}

	svc := service.NewRecommendationService(db)
	s := grpc.NewServer()
	pb.RegisterRecommendationServiceServer(s, handler.NewHandler(svc))

	log.Println("recommendation-service listening on :50053")
	log.Println("recommendation-service listening on :50053")
if err := s.Serve(lis); err != nil {
	log.Fatalf("gRPC server crashed: %v", err)
}
}
