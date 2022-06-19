package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	DBHost     string `env:"DB_HOST,notEmpty"`
	DBUser     string `env:"DB_USER,notEmpty"`
	DBPassword string `env:"DB_PASSWORD,notEmpty"`
	DBName     string `env:"DB_NAME,notEmpty"`
	DBPort     string `env:"DB_PORT,notEmpty"`
	JWTSecret  string `env:"JWT_SECRET,notEmpty"`
	HashCost   int    `env:"HASH_COST,notEmpty"`
	AppPort    string `env:"APP_PORT,notEmpty"`
}

func BuildConfig() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
