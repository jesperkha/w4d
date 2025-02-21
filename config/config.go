package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DbUrl string
}

func Load() (config Config, err error) {
	if err = godotenv.Load(); err != nil {
		return config, err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	if port[0] != ':' {
		port = ":" + port
	}

	config = Config{
		Port:  port,
		DbUrl: os.Getenv("DB_URL"),
	}

	return config, err
}
