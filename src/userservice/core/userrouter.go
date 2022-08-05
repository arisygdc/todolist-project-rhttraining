package core

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/pb"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/repository"
	"go-micro.dev/v4"
)

var ErrEmptyRequest = fmt.Errorf("empty request")

type Router struct {
	svc UserService
}

func NewRouter(svc UserService) Router {
	return Router{
		svc: svc,
	}
}

func (r Router) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest, empty *pb.Exec) error {
	err := r.svc.RegisterUser(ctx, UserDetail{
		Auth: repository.Auth{
			Username: request.Auth.Username,
			Password: request.Auth.Password,
			Email:    request.Auth.Email,
		},

		User: repository.User{
			FirstName: request.User.FirstName,
			LastName:  request.User.LastName,
		},
	})

	if err != nil {
		return err
	}

	empty.Success = true
	return nil
}

func (r Router) GetUser(ctx context.Context, request *pb.GetUserRequest, user *pb.User) error {
	if request == nil {
		return ErrEmptyRequest
	}

	id, err := uuid.Parse(request.UserId)
	if err != nil {
		return fmt.Errorf("not valid uuid")
	}

	getUser, err := r.svc.GetUser(ctx, id)
	if err != nil {
		return err
	}

	user.FirstName = getUser.FirstName
	user.LastName = getUser.LastName

	return nil
}

// RegisterHandler Register service and run server
func (r Router) RegisterHandler(service micro.Service) error {
	err := pb.RegisterUserServiceHandler(service.Server(), r)
	if err != nil {
		return err
	}

	service.Init()
	return service.Run()
}
