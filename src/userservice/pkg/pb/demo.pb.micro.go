// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: demo.proto

package pb

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...client.CallOption) (*Empty, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*User, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "UserService.RegisterUser", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUser", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	RegisterUser(context.Context, *RegisterUserRequest, *Empty) error
	GetUser(context.Context, *GetUserRequest, *User) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		RegisterUser(ctx context.Context, in *RegisterUserRequest, out *Empty) error
		GetUser(ctx context.Context, in *GetUserRequest, out *User) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) RegisterUser(ctx context.Context, in *RegisterUserRequest, out *Empty) error {
	return h.UserServiceHandler.RegisterUser(ctx, in, out)
}

func (h *userServiceHandler) GetUser(ctx context.Context, in *GetUserRequest, out *User) error {
	return h.UserServiceHandler.GetUser(ctx, in, out)
}

// Api Endpoints for AuthService service

func NewAuthServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AuthService service

type AuthService interface {
	AddAuth(ctx context.Context, in *Auth, opts ...client.CallOption) (*AddAuthResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) AddAuth(ctx context.Context, in *Auth, opts ...client.CallOption) (*AddAuthResponse, error) {
	req := c.c.NewRequest(c.name, "AuthService.AddAuth", in)
	out := new(AddAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceHandler interface {
	AddAuth(context.Context, *Auth, *AddAuthResponse) error
}

func RegisterAuthServiceHandler(s server.Server, hdlr AuthServiceHandler, opts ...server.HandlerOption) error {
	type authService interface {
		AddAuth(ctx context.Context, in *Auth, out *AddAuthResponse) error
	}
	type AuthService struct {
		authService
	}
	h := &authServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthService{h}, opts...))
}

type authServiceHandler struct {
	AuthServiceHandler
}

func (h *authServiceHandler) AddAuth(ctx context.Context, in *Auth, out *AddAuthResponse) error {
	return h.AuthServiceHandler.AddAuth(ctx, in, out)
}