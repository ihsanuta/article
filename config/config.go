package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig, MysqlConfig, RedisConfig map[string]interface{}

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	AppConfig = map[string]interface{}{
		"port": os.Getenv("SERVER_PORT"),
		"host": os.Getenv("SERVER_ADDRESS"),
	}

	MysqlConfig = map[string]interface{}{
		"username": os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASS"),
		"database": os.Getenv("DB_NAME"),
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
	}

	RedisConfig = map[string]interface{}{
		"host":     os.Getenv("REDIS_HOST"),
		"port":     os.Getenv("REDIS_PORT"),
		"database": os.Getenv("REDIS_DB"),
		"password": os.Getenv("REDIS_PASS"),
	}
}
