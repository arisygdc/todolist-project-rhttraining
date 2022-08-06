package main

import (
	"github.com/todolist-project-rhttraining/src/gateway/config"
	"github.com/todolist-project-rhttraining/src/gateway/core"
	"github.com/todolist-project-rhttraining/src/gateway/pkg/consumer"
	"go-micro.dev/v4"
	"log"
)

func main() {
	cfg := config.NewConfiguration()
	microService := micro.NewService()
	svc := consumer.NewConsumer(cfg.Endpoint, microService.Client())

	handler := core.NewHandler(svc)
	server := core.NewEchoServer(handler)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
