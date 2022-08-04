package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/todolist-project-rhttraining/src/userservice/config"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/pb"
	"go-micro.dev/v4/client"
	"sync"
)

type IDResponse struct {
	Id  uuid.UUID
	Err error
}

type PBRepo struct {
	Auth pb.AuthService
}

func NewProtoBuffRepo(serviceEndpoint config.ServiceEndpoint, c client.Client) PBRepo {
	authSvc := pb.NewAuthService(serviceEndpoint.Auth, c)

	return PBRepo{
		Auth: authSvc,
	}
}

func (pbRepo PBRepo) AddAuth(ctx context.Context, auth Auth, res *IDResponse, wait *sync.WaitGroup) {
	pbRes, err := pbRepo.Auth.AddAuth(ctx, &pb.Auth{
		Username: auth.Username,
		Password: auth.Password,
		Email:    auth.Email,
	})

	defer func() { wait.Done() }()

	if err != nil {
		res.Err = err
		return
	}

	authId, err := uuid.Parse(pbRes.GetId())
	if err != nil {
		res.Err = err
		return
	}

	res.Id = authId
	return
}
