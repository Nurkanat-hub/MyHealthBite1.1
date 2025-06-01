package handler

import (
	"context"
	"fmt"

	"email-service/internal/smtp"
	pb "email-service/proto"
)

type EmailHandler struct {
	pb.UnimplementedEmailServiceServer
}

func NewEmailHandler() *EmailHandler {
	return &EmailHandler{}
}

func (h *EmailHandler) SendEmail(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {
	err := smtp.Send(req.To, req.Subject, req.Body)
	if err != nil {
		return &pb.EmailResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to send email: %v", err),
		}, nil
	}

	return &pb.EmailResponse{
		Success: true,
		Message: "Email sent successfully",
	}, nil
}
