// user-service/internal/client/email_client.go
package client

import (
	pb "email-service/proto"
)

// EmailClient — глобальная переменная для доступа к gRPC-клиенту email-сервиса
var EmailClient pb.EmailServiceClient

// SetEmailClient — инициализирует email gRPC клиент
func SetEmailClient(c pb.EmailServiceClient) {
	EmailClient = c
}
