package config

import (
	"os"
)

const (
	endpointAuthKey = "SERVICE_AUTH_NAME"
	endpointUserKey = "SERVICE_USER_NAME"
	endpointTodoKey = "SERVICE_TODO_NAME"
)

type ServiceEndpoint struct {
	Auth string
	User string
	Todo string
}

func svcEndpointInit() ServiceEndpoint {
	if isEmpty(os.Getenv(endpointAuthKey)) {
		os.Setenv(endpointAuthKey, "svc-auth")
	}

	if isEmpty(os.Getenv(endpointUserKey)) {
		os.Setenv(endpointUserKey, "svc-user")
	}

	if isEmpty(os.Getenv(endpointTodoKey)) {
		os.Setenv(endpointTodoKey, "svc-todo")
	}

	return ServiceEndpoint{
		Auth: os.Getenv(endpointAuthKey),
		User: os.Getenv(endpointUserKey),
		Todo: os.Getenv(endpointTodoKey),
	}
}

func isEmpty(str string) bool {
	return str == ""
}
