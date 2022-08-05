package core

import "github.com/todolist-project-rhttraining/src/todoservice/pkg/repository"

type AuthService struct {
	repo repository.IRepository
}

func NewAuthService(repo repository.IRepository) AuthService {
	return AuthService{
		repo: repo,
	}
}
