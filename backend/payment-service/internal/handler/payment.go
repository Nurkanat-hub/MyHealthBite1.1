package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	emailpb "email-service/proto"
	"payment-service/internal/client"
	"payment-service/internal/repository"
	pb "payment-service/proto"
)

type PaymentHandler struct {
	pb.UnimplementedPaymentServiceServer
	Repo *repository.Repository
}

func NewPaymentHandler(repo *repository.Repository) *PaymentHandler {
	return &PaymentHandler{
		Repo: repo,
	}
}

func (h *PaymentHandler) MakePayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	status := "SUCCESS" // Принудительно успешный платёж

	payment := &repository.Payment{
		OrderID:  req.OrderId,
		UserName: req.UserName,
		Amount:   float64(req.Amount),
		Status:   status,
	}

	if err := h.Repo.Save(ctx, payment); err != nil {
		return &pb.PaymentResponse{
			Status:  "FAILURE",
			Message: "Failed to save payment: " + err.Error(),
		}, nil
	}

	// Асинхронная отправка письма при успехе
	if req.UserEmail != "" {
		go func(email, name, orderID string, amount float32) {
			emailCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			_, err := client.EmailClient.SendEmail(emailCtx, &emailpb.EmailRequest{
				To:      email,
				Subject: "Payment Confirmation",
				Body:    fmt.Sprintf("Hi %s,\n\nYour payment of $%.2f for order %s was successful.\n\nThanks for using MyHealthBite!", name, amount, orderID),
			})
			if err != nil {
				log.Printf("⚠️ Failed to send confirmation email to %s: %v", email, err)
			}
		}(req.UserEmail, req.UserName, req.OrderId, req.Amount)
	}

	return &pb.PaymentResponse{
		Status:  status,
		Message: "Payment processed",
	}, nil
}

func (h *PaymentHandler) GetPaymentHistory(ctx context.Context, req *pb.PaymentHistoryRequest) (*pb.PaymentHistoryResponse, error) {
	var fromDate, toDate time.Time
	var err error

	if req.FromDate != "" {
		fromDate, err = time.Parse(time.RFC3339, req.FromDate)
		if err != nil {
			return nil, err
		}
	}
	if req.ToDate != "" {
		toDate, err = time.Parse(time.RFC3339, req.ToDate)
		if err != nil {
			return nil, err
		}
	}

	filter := repository.PaymentFilter{
		UserName: req.UserName,
		Status:   req.Status,
		FromDate: fromDate,
		ToDate:   toDate,
		Limit:    int64(req.Limit),
	}

	records, err := h.Repo.GetHistory(ctx, filter)
	if err != nil {
		return nil, err
	}

	var paymentList []*pb.PaymentRecord
	for _, r := range records {
		paymentList = append(paymentList, &pb.PaymentRecord{
			OrderId:   r.OrderID,
			Amount:    r.Amount,
			Status:    r.Status,
			CreatedAt: r.CreatedAt.Format(time.RFC3339),
		})
	}

	return &pb.PaymentHistoryResponse{
		Payments: paymentList,
	}, nil
}
