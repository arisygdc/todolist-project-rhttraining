package core

import (
	"context"
	"github.com/todolist-project-rhttraining/src/authservice/pkg/pb"
	"go-micro.dev/v4"
)

type Router struct {
	svc AuthService
}

func NewRouter(svc AuthService) Router {
	return Router{
		svc: svc,
	}
}

func (r Router) AddAuth(ctx context.Context, req *pb.Auth, res *pb.AddAuthResponse) error {
	id, err := r.svc.AddAuth(ctx, req.Username, req.Password, req.Email)
	if err != nil {
		return err
	}

	res.Id = id
	return nil
}

func (r Router) Login(ctx context.Context, req *pb.LoginRequest, res *pb.Session) error {
	token, err := r.svc.Login(ctx, req.Username, req.Password)

	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (r Router) VerifyToken(ctx context.Context, req *pb.Session, res *pb.IdResponse) error {
	id, err := r.svc.VerifyToken(ctx, req.Token)
	if err != nil {
		return err
	}

	res.Id = id
	return nil
}

func (r Router) RegisterHandler(service micro.Service) error {
	err := pb.RegisterAuthServiceHandler(service.Server(), r)
	if err != nil {
		return err
	}

	service.Init()
	return service.Run()
}
