package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/internal/handler"
	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/internal/middleware"
	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/internal/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Загружаем .env переменные
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Основной роутер
	r := mux.NewRouter()

	// 🔓 Открытые маршруты
	r.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	r.HandleFunc("/register", handler.RegisterHandler).Methods("POST")

	// 🔐 Защищённые маршруты
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	// 🔐 Существующие защищённые
	protected.HandleFunc("/profile", handler.ProfileHandler).Methods("GET")
	protected.HandleFunc("/profile", handler.UpdateProfileHandler).Methods("PUT")
	protected.HandleFunc("/account", handler.DeleteAccountHandler).Methods("DELETE")
	protected.HandleFunc("/stats/{user_id}", handler.GetStatsHandler).Methods("GET")
	protected.HandleFunc("/stats/update", handler.UpdateStatsHandler).Methods("POST")
	protected.HandleFunc("/stats/reset", handler.ResetStatsHandler).Methods("POST")
	protected.HandleFunc("/stats/{user_id}", handler.DeleteStatsHandler).Methods("DELETE")

	// Подключение к user-service по gRPC

	connUser, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user-service: %v", err)
	}
	defer connUser.Close()

	userClient := service.NewUserClient(connUser)
	handler.SetUserClient(userClient)

	// Подключение к stats-service по gRPC
	connStats, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to stats-service: %v", err)
	}
	defer connStats.Close()

	statsClient := service.NewStatsClient(connStats)
	handler.SetStatsClient(statsClient)

	// ✅ Подключение к payment-service по gRPC
	connPayment, err := grpc.Dial("localhost:50057", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to payment-service: %v", err)
	}
	defer connPayment.Close()

	paymentClient := service.NewPaymentClient(connPayment)
	paymentHandler := handler.NewPaymentHandler(paymentClient)

	// ✅ Роуты оплаты
	protected.HandleFunc("/payment", paymentHandler.MakePayment).Methods("POST")
	protected.HandleFunc("/payment/history", paymentHandler.GetPaymentHistory).Methods("GET")

	// Запуск
	log.Printf("🚀 API Gateway is running on port %s", port)
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)
	if err := http.ListenAndServe(":"+port, corsHandler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
