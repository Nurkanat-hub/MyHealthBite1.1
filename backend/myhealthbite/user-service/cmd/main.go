package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"user-service/internal/client"
	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/proto"

	emailpb "email-service/proto"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Запускаем HTTP-сервер с метриками Prometheus
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("📊 Prometheus metrics available at :2112/metrics")
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatalf("Failed to start Prometheus metrics server: %v", err)
		}
	}()

	// Загрузка .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// Чтение порта из .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	// Чтение строки подключения к MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is required in .env file")
	}

	// Подключение к MongoDB
	clientMongo, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := clientMongo.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting MongoDB: %v", err)
		}
	}()

	userCollection := clientMongo.Database("myhealthbite_user").Collection("users")

	// ✅ Инициализация gRPC клиента для stats-service
	client.InitStatsClient()

	// ✅ Инициализация gRPC клиента для email-service
	emailAddr := os.Getenv("EMAIL_SERVICE_ADDR")
	if emailAddr == "" {
		log.Fatal("EMAIL_SERVICE_ADDR is required in .env file")
	}
	connEmail, err := grpc.Dial(emailAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to email-service: %v", err)
	}
	defer connEmail.Close()
	emailClient := emailpb.NewEmailServiceClient(connEmail)
	client.SetEmailClient(emailClient)

	// Запуск gRPC-сервера
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	// Создание и регистрация User-сервиса
	userRepo := repository.NewUserRepository(userCollection)
	userHandler := handler.NewUserServer(userRepo)
	proto.RegisterUserServiceServer(grpcServer, userHandler)

	// Включаем поддержку reflection для grpcurl и grpcui
	reflection.Register(grpcServer)

	log.Printf("✅ UserService is running on port %s...", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
