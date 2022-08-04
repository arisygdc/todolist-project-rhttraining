package config

import (
	"os"
)

type Configuration struct {
	Database    DBConfig
	SvcEndpoint ServiceEndpoint
}

type ServiceEndpoint struct {
	Auth string
	User string
}

type DBConfig struct {
	Host    string
	User    string
	Pass    string
	Port    string
	Name    string
	SslMode string
}

func NewConfiguration() Configuration {
	return Configuration{
		Database:    dbConfigurationInit(),
		SvcEndpoint: svcEndpointInit(),
	}
}

func svcEndpointInit() ServiceEndpoint {
	if isEmpty(os.Getenv(endpointAuthKey)) {
		os.Setenv(endpointAuthKey, "svc-auth")
	}

	if isEmpty(os.Getenv(endpointUserKey)) {
		os.Setenv(endpointUserKey, "svc-user")
	}

	return ServiceEndpoint{
		Auth: os.Getenv(endpointAuthKey),
		User: os.Getenv(endpointUserKey),
	}
}

func dbConfigurationInit() DBConfig {
	if isEmpty(os.Getenv(dbHostKey)) {
		os.Setenv(dbHostKey, "127.0.0.1")
	}

	if isEmpty(os.Getenv(dbUserKey)) {
		os.Setenv(dbUserKey, "rhtPostgreUser")
	}

	if isEmpty(os.Getenv(dbPassKey)) {
		os.Setenv(dbPassKey, "rhtPostgrePassword")
	}

	if isEmpty(os.Getenv(dbPortKey)) {
		os.Setenv(dbPortKey, "5432")
	}

	if isEmpty(os.Getenv(dbNameKey)) {
		os.Setenv(dbNameKey, "user")
	}

	if isEmpty(os.Getenv(dbSSLModeKey)) || (os.Getenv(dbSSLModeKey) != "disable" && os.Getenv(dbSSLModeKey) != "enable") {
		os.Setenv(dbSSLModeKey, "disable")
	}

	return DBConfig{
		Host:    os.Getenv(dbHostKey),
		User:    os.Getenv(dbUserKey),
		Pass:    os.Getenv(dbPassKey),
		Port:    os.Getenv(dbPortKey),
		Name:    os.Getenv(dbNameKey),
		SslMode: os.Getenv(dbSSLModeKey),
	}
}

func isEmpty(str string) bool {
	return str == ""
}
