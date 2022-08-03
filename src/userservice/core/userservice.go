package core

import (
	"context"
	"github.com/google/uuid"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/repository"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/repository/db"
)

type UserDetail struct {
	repository.Auth
	repository.User
}

type UserService struct {
	dbRepo repository.DatabaseRepo
	pbRepo repository.PBRepo
}

func NewUserService(dbRepo repository.DatabaseRepo, pbRepo repository.PBRepo) UserService {
	return UserService{
		dbRepo: dbRepo,
		pbRepo: pbRepo,
	}
}

func (us UserService) RegisterUser(ctx context.Context, detail UserDetail) error {
	ch := make(chan repository.IDResponse)
	defer close(ch)

	go us.pbRepo.AddAuth(ctx, detail.Auth, ch)

	param := db.AddUserParams{
		ID:        uuid.New(),
		FirstName: detail.FirstName,
		LastName:  detail.LastName,
	}

	res := <-ch
	if res.Err != nil {
		return res.Err
	}

	param.AuthID = res.Id

	return us.dbRepo.Query().AddUser(ctx, param)
}

func (us UserService) GetUser(ctx context.Context, id uuid.UUID) (db.User, error) {
	return us.dbRepo.Query().GetUser(ctx, id)
}
