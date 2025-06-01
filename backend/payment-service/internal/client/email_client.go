package client

import (
	pb "email-service/proto"

	"google.golang.org/grpc"
)

var EmailClient pb.EmailServiceClient

func InitEmailClient(cc *grpc.ClientConn) {
	EmailClient = pb.NewEmailServiceClient(cc)
}
