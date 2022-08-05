package core

import (
	"context"
	"github.com/google/uuid"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/repository"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/repository/db"
	"sync"
)

var wg = new(sync.WaitGroup)

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
	res := repository.IDResponse{}

	wg.Add(1)
	go us.pbRepo.AddAuth(ctx, detail.Auth, &res, wg)

	param := db.AddUserParams{
		ID:        uuid.New(),
		FirstName: detail.FirstName,
		LastName:  detail.LastName,
	}

	wg.Wait()

	if res.Err != nil {
		return res.Err
	}

	param.AuthID = res.Id

	return us.dbRepo.Query().AddUser(ctx, param)
}

func (us UserService) GetUser(ctx context.Context, id uuid.UUID) (db.User, error) {
	return us.dbRepo.Query().GetUser(ctx, id)
}
