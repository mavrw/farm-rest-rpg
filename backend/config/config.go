package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type AuthConfig struct {
	JWTSecret string
}

type ServerConfig struct {
	Port string
}

type Config struct {
	Environment string
	DB          DBConfig
	Auth        AuthConfig
	Server      ServerConfig
}

func Load() (Config, error) {
	_ = godotenv.Load("devops/.env")

	return Config{
		Environment: os.Getenv("ENV"),
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Auth: AuthConfig{
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
		Server: ServerConfig{
			Port: os.Getenv("API_PORT"),
		},
	}, nil
}
