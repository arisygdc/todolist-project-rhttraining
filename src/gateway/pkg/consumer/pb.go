package consumer

import (
	"github.com/todolist-project-rhttraining/src/gateway/config"
	"github.com/todolist-project-rhttraining/src/gateway/pkg/pb"
	"go-micro.dev/v4/client"
)

type Consumer struct {
	User pb.UserService
	Auth pb.AuthService
	Todo pb.TodoService
}

func NewConsumer(endpoint config.ServiceEndpoint, c client.Client) Consumer {
	return Consumer{
		User: pb.NewUserService(endpoint.User, c),
		Auth: pb.NewAuthService(endpoint.Auth, c),
		Todo: pb.NewTodoService(endpoint.Todo, c),
	}
}
