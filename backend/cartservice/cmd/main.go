package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"cartservice/internal/repository"
	"cartservice/internal/service"
	pb "cartservice/proto"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")
	dbName := "myhealthbite"
	collectionName := "cart"

	// MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}
	collection := client.Database(dbName).Collection(collectionName)

	// Repository
	cartRepo := repository.NewCartRepository(collection)

	// gRPC server
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterCartServiceServer(grpcServer, service.NewCartService(cartRepo))

	fmt.Println("Cart gRPC server is running on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
