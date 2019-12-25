package auth

import (
	"os"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	HTTPPort        string `validate:"required"`
	LogVerbose      bool
	UserServiceHost string `validate:"required"`
	JWTSecret       string `validate:"required"`
	DB              DBConfig
}

type DBConfig struct {
	Host     string `validate:"required"`
	Port     string `validate:"required,numeric"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
	SSL      string `validate:"required,oneof=enable disable"`
}

// nolint: funlen
func InitConfig() (*Config, error) {
	v := viper.New()
	v.AutomaticEnv()

	c := Config{
		HTTPPort:        v.GetString("HTTP_PORT"),
		LogVerbose:      v.GetBool("LOG_VERBOSE"),
		UserServiceHost: v.GetString("USER_SERVICE_HOST"),
		JWTSecret:       v.GetString("JWT_SECRET"),
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Database: os.Getenv("DB_SCHEMA"),
			SSL:      os.Getenv("DB_SSL"),
		},
	}

	if err := validator.New().Struct(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
