package service

import (
	"github.com/Nurkanat-hub/MyHealthbite/api-gateway/proto"
	"google.golang.org/grpc"
)

type UserClient struct {
	Client proto.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		Client: proto.NewUserServiceClient(conn),
	}
}
