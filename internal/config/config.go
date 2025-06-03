package config

import (
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/scdb/core/logger"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		logger.Warning(".env файл не найден, пробую использовать системные переменные")
	}

	AppConfig = &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "db"),
	}

	logger.Success("Конфигурация env загружена")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
