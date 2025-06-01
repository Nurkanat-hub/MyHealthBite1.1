package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/internal/service"
	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/proto"
	"github.com/golang-jwt/jwt/v5"
)

var userClient *service.UserClient

func SetUserClient(c *service.UserClient) {
	userClient = c
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req proto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := userClient.Client.Login(ctx, &req)
	if err != nil {
		http.Error(w, "Login failed: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// ✅ Создание JWT токена с user_id и user_name
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		http.Error(w, "JWT secret is missing", http.StatusInternalServerError)
		return
	}

	claims := jwt.MapClaims{
		"user_id":   resp.UserId,
		"user_name": resp.Name, // ВАЖНО: это поле должно быть возвращено user-service
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		http.Error(w, "Failed to sign token", http.StatusInternalServerError)
		return
	}

	// ✅ Возвращаем токен и userId
	json.NewEncoder(w).Encode(map[string]string{
		"token":    tokenStr,
		"userId":   resp.UserId,
		"userName": resp.Name, // для отладки можно вернуть и имя
	})
}
