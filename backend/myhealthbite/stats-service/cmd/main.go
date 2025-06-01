package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"MyHealthBite/stats-service/internal/handler"
	"MyHealthBite/stats-service/internal/repository"
	"MyHealthBite/stats-service/proto"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "50054"
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("Missing MONGO_URI in .env")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}
	defer client.Disconnect(context.Background())

	statsCollection := client.Database("myhealthbite_stats").Collection("stats")

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	repo := repository.NewStatsRepository(statsCollection)
	statsServer := handler.NewStatsServer(repo)

	proto.RegisterStatsServiceServer(grpcServer, statsServer)
	reflection.Register(grpcServer)

	log.Printf("StatsService is running on port %s...", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC Server error: %v", err)
	}
}
