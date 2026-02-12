package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Env string `env:"ENV" envDefault:"local"`
	//

	BotToken string `env:"BOT_TOKEN,required"`

	//
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBSSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

func ReadEnvConfig(cfg any) error {
	return env.Parse(cfg)
}
