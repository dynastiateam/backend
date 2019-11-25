package config

import (
	"os"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	HTTPPort  string `validate:"required,numeric"`
	JWTSecret string `validate:"required"`
	Debug     bool
	DB        DBConfig
}

type DBConfig struct {
	Host     string `validate:"required"`
	Port     string `validate:"required,numeric"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
	SSL      string `validate:"required,oneof=enable disable"`
}

func New() (*Config, error) {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return nil, err
	}

	c := &Config{
		HTTPPort:  os.Getenv("HTTP_PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		Debug:     debug,
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Database: os.Getenv("DB_SCHEMA"),
			SSL:      os.Getenv("DB_SSL"),
		},
	}

	if err := validator.New().Struct(c); err != nil {
		return nil, err
	}

	return c, nil
}
