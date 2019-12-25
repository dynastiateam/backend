package auth

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

func (mw loggingMiddleware) Login(ctx context.Context, req *loginRequest) (*loginResponse, error) {
	res, err := mw.next.Login(ctx, req)
	if err != nil {
		mw.log.Error().Msg(err.Error())
	}

	return res, err
}

func (mw loggingMiddleware) Refresh(ctx context.Context, req *refreshTokenRequest) (*loginResponse, error) {
	res, err := mw.next.Refresh(ctx, req)
	if err != nil {
		mw.log.Error().Msg(err.Error())
	}

	return res, err
}
