package config

import "os"

const (
	dbDriverKey   = "DB_DRIVER"
	dbHostKey     = "DB_HOST"
	dbUserKey     = "DB_USER"
	dbPasswordKey = "DB_PASSWORD"
	dbNameKey     = "DB_NAME"

	endpointTodoKey = "SERVICE_TODO_NAME"
)

type Configuration struct {
	Database DbConfig
	Endpoint ServiceEndpoint
}

type ServiceEndpoint struct {
	Todo string
}

type DbConfig struct {
	Driver   string
	Host     string
	User     string
	Password string
	DbName   string
}

func NewConfiguration() Configuration {
	return Configuration{
		Database: dbConfigInit(),
		Endpoint: svcEndpointInit(),
	}
}

func svcEndpointInit() ServiceEndpoint {
	if isEmpty(os.Getenv(endpointTodoKey)) {
		os.Setenv(endpointTodoKey, "svc-todo")
	}

	return ServiceEndpoint{
		Todo: os.Getenv(endpointTodoKey),
	}
}

func dbConfigInit() DbConfig {
	if isEmpty(os.Getenv(dbDriverKey)) {
		os.Setenv(dbDriverKey, "mongodb")
	}

	if isEmpty(os.Getenv(dbHostKey)) {
		os.Setenv(dbHostKey, "127.0.0.1")
	}

	if isEmpty(os.Getenv(dbUserKey)) {
		os.Setenv(dbUserKey, "rhtMongoUser")
	}

	if isEmpty(os.Getenv(dbPasswordKey)) {
		os.Setenv(dbPasswordKey, "rhtMongoPassword")
	}

	if isEmpty(os.Getenv(dbNameKey)) {
		os.Setenv(dbNameKey, "todo")
	}

	return DbConfig{
		Driver:   os.Getenv(dbDriverKey),
		Host:     os.Getenv(dbHostKey),
		User:     os.Getenv(dbUserKey),
		Password: os.Getenv(dbPasswordKey),
		DbName:   os.Getenv(dbNameKey),
	}
}

func isEmpty(str string) bool {
	return str == ""
}
