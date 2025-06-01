package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/proto"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req proto.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	resp, err := userClient.Client.Register(ctx, &req)
	if err != nil {
		// Извлекаем только читаемое сообщение
		cleanMsg := cleanGrpcError(err.Error())

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": cleanMsg,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// cleanGrpcError удаляет "rpc error: code = ..." и возвращает только сообщение
func cleanGrpcError(raw string) string {
	// Пример: "rpc error: code = AlreadyExists desc = email already registered"
	if i := strings.Index(raw, "desc = "); i != -1 {
		return raw[i+7:]
	}
	return raw
}
