package core

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	handler  IHandler
	provider *echo.Echo
}

func NewEchoServer(handler IHandler) EchoServer {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return EchoServer{
		handler:  handler,
		provider: e,
	}
}

func (e EchoServer) route() {
	e.provider.POST("/api/v1/register", e.handler.Register)
	e.provider.POST("/api/v1/login", e.handler.Login)

	authGroup := e.provider.Group("/api/v1/user", AuthenticatedUser)
	authGroup.GET("/name", e.handler.GetUser)
}

func (e EchoServer) Run() error {
	e.route()
	err := e.provider.Start(fmt.Sprintf(":8080"))
	e.provider.Logger.Fatal(err)
	return err
}
