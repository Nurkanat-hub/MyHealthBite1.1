package handler

import (
	"encoding/json"
	"log"
	"net/http"

	statsProto "MyHealthBite/stats-service/proto"

	"github.com/gorilla/mux"
)

// Переменная для хранения клиента stats-service
var statsClient statsProto.StatsServiceClient

// SetStatsClient устанавливает gRPC клиента для stats-service
func SetStatsClient(client statsProto.StatsServiceClient) {
	statsClient = client
}

// GetStatsHandler обрабатывает запрос GET /api/stats/{user_id}
func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	if statsClient == nil {
		http.Error(w, "Stats service client not initialized", http.StatusInternalServerError)
		return
	}

	resp, err := statsClient.GetStats(r.Context(), &statsProto.UserIdRequest{UserId: userID})
	if err != nil {
		log.Printf("Error calling GetStats: %v", err)
		http.Error(w, "Failed to get stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// UpdateStatsHandler обрабатывает POST /api/stats/update
func UpdateStatsHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID        string `json:"user_id"`
		DeltaCalories int32  `json:"delta_calories"`
		DeltaWaterML  int32  `json:"delta_water_ml"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := statsClient.UpdateStats(r.Context(), &statsProto.UpdateStatsRequest{
		UserId:        req.UserID,
		DeltaCalories: req.DeltaCalories,
		DeltaWaterMl:  req.DeltaWaterML,
	})
	if err != nil {
		log.Printf("Error calling UpdateStats: %v", err)
		http.Error(w, "Failed to update stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ResetStatsHandler обрабатывает POST /api/stats/reset
func ResetStatsHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID string `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := statsClient.ResetDailyStats(r.Context(), &statsProto.UserIdRequest{UserId: req.UserID})
	if err != nil {
		log.Printf("Error calling ResetDailyStats: %v", err)
		http.Error(w, "Failed to reset stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DeleteStatsHandler обрабатывает DELETE /api/stats/{user_id}
func DeleteStatsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	_, err := statsClient.DeleteStatsByUserId(r.Context(), &statsProto.UserIdRequest{UserId: userID})
	if err != nil {
		log.Printf("Error calling DeleteStatsByUserId: %v", err)
		http.Error(w, "Failed to delete stats", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
