package config

import (
	"log"
	"os"
)

type Config struct {
	ServerPort string
	JWTSecret  string
}

func LoadConfig() Config {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	return Config{
		ServerPort: port,
		JWTSecret:  jwtSecret,
	}
}
