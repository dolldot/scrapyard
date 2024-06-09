package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/dolldot/scrapyard/bookmarkd/internal/logger"
	"github.com/joho/godotenv"
)

var Load = struct {
	DB struct {
		Type     string `env:"DB_TYPE" envDefault:"postgres"`
		Host     string `env:"DB_HOST" envDefault:"127.0.0.1"`
		Port     string `env:"DB_PORT" envDefault:"5432"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Name     string `env:"DB_NAME" envDefault:"bookmarkd"`
	}

	Server struct {
		Host     string `env:"HOST" envDefault:"localhost"`
		Port     string `env:"PORT" envDefault:":8080"`
		Mode     string `env:"MODE" envDefault:"debug"`
		TimeZone string `env:"TIMEZONE" envDefault:"Asia/Jakarta"`
	}
}{}

func InitializeConfig() {
	godotenv.Load()

	err := env.Parse(&Load)
	if err != nil {
		logger.Fatalf("Unable to parse environment variables: %e", err)
	}
}
