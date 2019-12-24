package user

import (
	"context"

	"github.com/rs/zerolog"
)

type loggingMiddleware struct {
	log  *zerolog.Logger
	next Service
}

func newLoggingMiddleware(log *zerolog.Logger, svc Service) Service {
	return loggingMiddleware{
		log:  log,
		next: svc,
	}
}

func (mw loggingMiddleware) Hello(ctx context.Context, key string) (interface{}, error) {
	res, err := mw.next.Hello(ctx, key)
	if err != nil {
		mw.log.Error().Msg(err.Error())
	}

	return res, err
}
