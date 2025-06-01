package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/internal/middleware"
	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/proto"
)

// Получение профиля по userId из токена
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
		return
	}

	// Debug log userID extracted from token
	println("ProfileHandler: userID from token:", userID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := userClient.Client.GetProfile(ctx, &proto.UserIdRequest{UserId: userID})
	if err != nil {
		http.Error(w, "Failed to get profile: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Обновление профиля (goal, height, phone и т.д.)
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req proto.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	req.UserId = userID // добавляем userId из JWT

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := userClient.Client.UpdateProfile(ctx, &req)
	if err != nil {
		http.Error(w, "Failed to update profile: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
