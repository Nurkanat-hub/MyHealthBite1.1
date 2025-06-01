package service

import (
	"context"
	"fmt"
	"time"

	pb "payment-service/proto"

	"google.golang.org/grpc"
)

type PaymentClient struct {
	client pb.PaymentServiceClient
}

func NewPaymentClient(conn *grpc.ClientConn) *PaymentClient {
	return &PaymentClient{
		client: pb.NewPaymentServiceClient(conn),
	}
}

func (p *PaymentClient) MakePayment(ctx context.Context, orderID string, userName string, userEmail string, amount float32) (*pb.PaymentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := p.client.MakePayment(ctx, &pb.PaymentRequest{
		OrderId:   orderID,
		UserName:  userName,
		UserEmail: userEmail,
		Amount:    amount,
	})
	if err != nil {
		return nil, fmt.Errorf("MakePayment failed: %w", err)
	}
	return res, nil
}

func (p *PaymentClient) GetPaymentHistory(ctx context.Context, userName string) (*pb.PaymentHistoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := p.client.GetPaymentHistory(ctx, &pb.PaymentHistoryRequest{
		UserName: userName,
	})
	if err != nil {
		return nil, fmt.Errorf("GetPaymentHistory failed: %w", err)
	}
	return res, nil
}
