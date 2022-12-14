package main

import (
	"context"
	"github.com/todolist-project-rhttraining/src/userservice/config"
	"github.com/todolist-project-rhttraining/src/userservice/core"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/repository"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.NewConfiguration()

	mServer := micro.NewService(
		micro.Name(cfg.SvcEndpoint.User),
	)

	dbRepo, err := repository.NewDatabaseRepo(ctx, cfg.Database)
	if err != nil {
		log.Panic(err)
	}

	cln := client.NewClient()
	pbRepo := repository.NewProtoBuffRepo(cfg.SvcEndpoint, cln)
	userService := core.NewUserService(dbRepo, pbRepo)

	r := core.NewRouter(userService)
	err = r.RegisterHandler(mServer)
	if err != nil {
		log.Fatal(err)
	}
}
