package main

import (
	"github.com/todolist-project-rhttraining/src/todoservice/config"
	"github.com/todolist-project-rhttraining/src/todoservice/core"
	"github.com/todolist-project-rhttraining/src/todoservice/pkg/pb"
	"github.com/todolist-project-rhttraining/src/todoservice/pkg/repository"
	"go-micro.dev/v4"
	"log"
)

func main() {
	cfg := config.NewConfiguration()

	svcProvider := micro.NewService(
		micro.Name(cfg.Endpoint.Todo),
	)

	repo, err := repository.NewRepository(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	svc := core.NewTodoService(repo)
	handler := core.NewHandler(svc)
	err = pb.RegisterTodoServiceHandler(svcProvider.Server(), handler)
	if err != nil {
		log.Fatal(err)
	}

	svcProvider.Init()
	if err := svcProvider.Run(); err != nil {
		log.Fatal(err)
	}
}
