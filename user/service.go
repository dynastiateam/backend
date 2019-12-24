package user

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
)

type Service interface {
	Hello(ctx context.Context, key string) (interface{}, error)
}

type service struct {
}

func NewService(log *zerolog.Logger) Service {
	s := &service{}
	svc := newLoggingMiddleware(log, s)

	return svc
}

func (s *service) Hello(_ context.Context, key string) (interface{}, error) {
	if key == "1" {
		return nil, errors.New("some error")
	}
	return "service is working", nil
}
