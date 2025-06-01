package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/internal/middleware"
	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/proto"
)

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем userId из JWT
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userClient.Client.DeleteAccount(ctx, &proto.UserIdRequest{UserId: userID})
	if err != nil {
		http.Error(w, "Failed to delete account: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
