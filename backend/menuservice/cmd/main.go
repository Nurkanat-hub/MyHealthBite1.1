package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"menuservice/internal/repository"
	"menuservice/internal/service"
	proto "menuservice/proto"
)

func main() {
	// Загрузка .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")
	if port == "" {
		port = "50059"
	}

	// Подключение к MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collection := client.Database("myhealthbite_menu").Collection("dishes")
	repository.SetCollection(collection) // передаём Mongo внутрь репозитория

	// Настройка gRPC
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	repo := repository.NewMongoRepository()
	proto.RegisterMenuServiceServer(grpcServer, service.NewMenuService(repo))

	fmt.Println("🚀 gRPC server started on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
