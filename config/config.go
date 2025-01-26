package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() (config Config, err error) {
	if err = godotenv.Load(); err != nil {
		return config, err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config = Config{
		Port: port,
	}

	return config, err
}
