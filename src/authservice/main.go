package main

import (
	"context"
	"github.com/todolist-project-rhttraining/src/authservice/config"
	"github.com/todolist-project-rhttraining/src/authservice/core"
	"github.com/todolist-project-rhttraining/src/authservice/pkg/repository"
	"github.com/todolist-project-rhttraining/src/authservice/pkg/token"
	"go-micro.dev/v4"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.NewConfiguration()

	mServer := micro.NewService(
		micro.Name(cfg.SvcEndpoint.Auth),
	)

	dbRepo, err := repository.NewDatabaseRepo(ctx, cfg.Database)
	if err != nil {
		log.Panic(err)
	}

	tokenMaker := token.NewJWT(cfg.Token.Secret)

	authService := core.NewAuthService(dbRepo, tokenMaker)

	r := core.NewRouter(authService)
	err = r.RegisterHandler(mServer)
	if err != nil {
		log.Fatal(err)
	}
}
