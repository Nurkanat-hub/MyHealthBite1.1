package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	emailpb "email-service/proto"
	"payment-service/internal/client"
	"payment-service/internal/handler"
	"payment-service/internal/repository"
	pb "payment-service/proto"
)

func main() {
	// Загрузка .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env not found, using system environment")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("❌ MONGO_URI is not set")
	}

	clientMongo, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("❌ Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := clientMongo.Connect(ctx); err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	db := clientMongo.Database("myhealthdite_payment")
	collection := db.Collection("payment")

	repo := repository.NewRepository(collection)
	paymentHandler := handler.NewPaymentHandler(repo)

	// Подключение к email-service
	emailAddr := os.Getenv("EMAIL_SERVICE_ADDR")
	if emailAddr == "" {
		emailAddr = "localhost:50058"
	}
	connEmail, err := grpc.Dial(emailAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Failed to connect to email-service: %v", err)
	}
	defer connEmail.Close()

	client.EmailClient = emailpb.NewEmailServiceClient(connEmail)

	// Запуск gRPC сервера
	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, paymentHandler)

	// Включаем reflection для grpcurl
	reflection.Register(grpcServer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "50057"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("❌ Failed to listen on port %s: %v", port, err)
	}

	fmt.Printf("✅ Payment Service running on port %s\n", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Failed to serve gRPC: %v", err)
	}
}
