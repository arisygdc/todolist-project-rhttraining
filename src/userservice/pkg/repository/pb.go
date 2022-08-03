package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/todolist-project-rhttraining/src/userservice/config"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/pb"
	"go-micro.dev/v4/client"
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

func (pbRepo PBRepo) AddAuth(ctx context.Context, auth Auth, ch chan IDResponse) {
	var res = IDResponse{
		Id:  uuid.UUID{},
		Err: nil,
	}
	pbRes, err := pbRepo.Auth.AddAuth(ctx, &pb.Auth{
		Username: auth.Username,
		Password: auth.Password,
		Email:    auth.Email,
	})

	if err != nil {
		res.Err = err
		ch <- res
		return
	}

	authId, err := uuid.Parse(pbRes.GetId())
	if err != nil {
		res.Err = err
		ch <- res
		return
	}

	res.Id = authId
	ch <- res
}
