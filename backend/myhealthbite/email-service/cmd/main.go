// cmd/main.go
package main

import (
	"log"
	"net"
	"os"

	"email-service/internal/handler"
	pb "email-service/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Загрузка .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	port := os.Getenv("EMAIL_SERVICE_PORT")
	if port == "" {
		port = "50058"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, handler.NewEmailHandler())
	reflection.Register(grpcServer)

	log.Printf("📨 Email Service running on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
