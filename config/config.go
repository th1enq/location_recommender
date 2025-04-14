package config

import (
	"github.com/joeshaw/envdecode"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port string `env:"SERVER_PORT,required"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Name     string `env:"DB_NAME,required"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
