package configs

import (
	"os"
	"strconv"
)

// AppConfig - структура настроек
type AppConfig struct {
	AppPort     string
	DBUri       string
	TokenLength int
}

// ProjectSettings - глобальные настройки проекта
var ProjectSettings AppConfig

// Settings - основные настройки сервиса
func Settings() AppConfig {
	var appPort = os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8000"
	}
	var dbURI = os.Getenv("DB_URI")
	if dbURI == "" {
		dbURI = "mongodb://root:example@localhost:27017"
	}
	tokenLength, err := strconv.Atoi(os.Getenv("TOKEN_LENGTH"))
	if err != nil {
		tokenLength = 6
	}

	ProjectSettings = AppConfig{
		AppPort:     appPort,
		DBUri:       dbURI,
		TokenLength: tokenLength,
	}

	return ProjectSettings
}
