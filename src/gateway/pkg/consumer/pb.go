package consumer

import (
	"github.com/todolist-project-rhttraining/src/gateway/config"
	"github.com/todolist-project-rhttraining/src/gateway/pkg/pb"
	"go-micro.dev/v4/client"
)

type Consumer struct {
	User pb.UserService
	Auth pb.AuthService
}

func NewConsumer(endpoint config.ServiceEndpoint) Consumer {
	return Consumer{
		User: pb.NewUserService(endpoint.User, client.NewClient()),
		Auth: pb.NewAuthService(endpoint.Auth, client.NewClient()),
	}
}
