package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/internal/middleware"
	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/internal/service"
)

type PaymentHandler struct {
	PaymentClient *service.PaymentClient
}

func NewPaymentHandler(client *service.PaymentClient) *PaymentHandler {
	return &PaymentHandler{PaymentClient: client}
}

// POST /payment
func (h *PaymentHandler) MakePayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrderID   string  `json:"order_id"`
		Amount    float32 `json:"amount"`
		UserEmail string  `json:"user_email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userName := middleware.GetUserName(r.Context())
	if userName == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	res, err := h.PaymentClient.MakePayment(r.Context(), req.OrderID, userName, req.UserEmail, req.Amount)
	if err != nil {
		http.Error(w, "Payment failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// GET /payment/history
func (h *PaymentHandler) GetPaymentHistory(w http.ResponseWriter, r *http.Request) {
	userName := middleware.GetUserName(r.Context())
	if userName == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	res, err := h.PaymentClient.GetPaymentHistory(r.Context(), userName)
	if err != nil {
		http.Error(w, "Failed to fetch history: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
