package core

import (
	"context"
	"github.com/google/uuid"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/repository"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/repository/db"
	"log"
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
	log.Println("--- Start Request RegisterUser")
	log.Printf("param: %v\n", detail)

	defer func() { log.Println("--- End Request RegisterUser") }()

	res := repository.IDResponse{}

	log.Printf("send %v to svc auth\n", detail.Auth)
	wg.Add(1)
	go us.pbRepo.AddAuth(ctx, detail.Auth, &res, wg)

	param := db.AddUserParams{
		ID:        uuid.New(),
		FirstName: detail.FirstName,
		LastName:  detail.LastName,
	}

	wg.Wait()

	if res.Err != nil {
		log.Printf("failed with error: %v\n", res.Err)
		return res.Err
	}

	param.AuthID = res.Id

	return us.dbRepo.Query().AddUser(ctx, param)
}

func (us UserService) GetUser(ctx context.Context, id uuid.UUID) (db.User, error) {
	log.Println("--- Start Request GetUser")
	log.Printf("id: %v\n", id)

	defer func() { log.Println("--- End Request GetUser") }()

	log.Printf("geting user from id: %s\n", id.String())
	usr, err := us.dbRepo.Query().GetUser(ctx, id)
	if err != nil {
		log.Printf("failed with error: %v\n", err)
	}

	log.Printf("found: %v\n", usr)
	return usr, err
}
