package core

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/todolist-project-rhttraining/src/gateway/pkg/consumer"
	"github.com/todolist-project-rhttraining/src/gateway/pkg/pb"
	"net/http"
	"time"
)

type IHandler interface {
	Register(ctx echo.Context) (err error)
	Login(ctx echo.Context) (err error)
	GetUser(c echo.Context) (err error)
}

type Handler struct {
	svc consumer.Consumer
}

func NewHandler(svc consumer.Consumer) Handler {
	return Handler{svc: svc}
}

func (h Handler) Register(c echo.Context) (err error) {
	req := new(RegisterUser)
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		return ResponseBadRequest(c, err.Error())
	}

	if err := ValidateRequest(req); err != nil {
		return ResponseBadRequest(c, err.Error())
	}

	param := pb.RegisterUserRequest{
		Auth: &pb.Auth{
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
		},
		User: &pb.User{
			FirstName: req.Firstname,
			LastName:  req.Lastname,
		},
	}
	execStatus, err := h.svc.User.RegisterUser(ctx, &param)
	if err != nil {
		return ResponseBadRequest(c, err.Error())
	}

	if !execStatus.Success {
		return ResponseBadRequest(c, "register gagal")
	}

	return ResponseCreated(c, map[string]string{
		"username": param.Auth.Username,
		"name":     fmt.Sprintf("%s %s", param.User.FirstName, param.User.LastName),
	})
}

func (h Handler) Login(c echo.Context) (err error) {
	req := new(Authentication)
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		return ResponseBadRequest(c, err.Error())
	}

	if err := ValidateRequest(req); err != nil {
		return ResponseBadRequest(c, err.Error())
	}

	param := pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	sessionToken, err := h.svc.Auth.Login(ctx, &param)
	if err != nil {
		return ResponseBadRequest(c, err.Error())
	}

	cookie := http.Cookie{
		Name:     "X-SESSION-TOKEN",
		Value:    sessionToken.Token,
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
	}

	c.SetCookie(&cookie)
	return ResponseOk(c, map[string]string{"token": sessionToken.Token})
}

func (h Handler) GetUser(c echo.Context) (err error) {
	id, err := validateToken(c, h.svc.Auth)

	if err != nil {
		return ResponseBadRequest(c, err.Error())
	}

	ctx := c.Request().Context()
	user, err := h.svc.User.GetUser(ctx, &pb.GetUserRequest{
		UserId: id,
	})

	if err != nil {
		return ResponseBadRequest(c, err.Error())
	}

	return ResponseOk(c, map[string]string{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	})
}
