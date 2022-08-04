package config

import (
	"os"
	"strconv"
)

type Configuration struct {
	Database    DBConfig
	SvcEndpoint ServiceEndpoint
}

type ServiceEndpoint struct {
	Auth string
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

	return ServiceEndpoint{
		Auth: os.Getenv(endpointAuthKey),
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
		os.Setenv(dbNameKey, "auth")
	}

	_, err := strconv.ParseBool(os.Getenv(dbSSLModeKey))

	if err != nil {
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
