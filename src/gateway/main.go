package main

import (
	"github.com/todolist-project-rhttraining/src/gateway/config"
	"github.com/todolist-project-rhttraining/src/gateway/core"
	"github.com/todolist-project-rhttraining/src/gateway/pkg/consumer"
	"log"
)

func main() {
	cfg := config.NewConfiguration()

	svc := consumer.NewConsumer(cfg.Endpoint)
	handler := core.NewHandler(svc)
	server := core.NewEchoServer(handler)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
