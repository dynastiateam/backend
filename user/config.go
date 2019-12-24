package user

import (
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	HTTPPort            string
	HTTPSPort           string `validate:"required_without=HTTPPort"`
	CORSWhiteListedHost string `validate:"required"`
	LogVerbose          bool
}

// nolint: funlen
func InitConfig(env string) (*Config, error) {
	v := viper.New()
	v.AutomaticEnv()

	c := Config{
		HTTPPort:            v.GetString("HTTP_PORT"),
		HTTPSPort:           v.GetString("HTTPS_PORT"),
		CORSWhiteListedHost: v.GetString("CORS_WHITELISTED_HOST"),
		LogVerbose:          v.GetBool("LOG_VERBOSE"),
	}

	if err := validator.New().Struct(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
