package users

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

func (mw loggingMiddleware) Register(ctx context.Context, req *userRegisterRequest) (*userRegisterResponse, error) {
	res, err := mw.next.Register(ctx, req)
	if err != nil {
		mw.log.Error().Msg(err.Error())
	}

	return res, err
}

func (mw loggingMiddleware) UserByEmailAndPassword(ctx context.Context, r *userByEmailAndPasswordRequest) (*User, error) {
	res, err := mw.next.UserByEmailAndPassword(ctx, r)
	if err != nil {
		mw.log.Error().Msg(err.Error())
	}

	return res, err
}

func (mw loggingMiddleware) UserByID(ctx context.Context, id int) (*userByIDResponse, error) {
	res, err := mw.next.UserByID(ctx, id)
	if err != nil {
		mw.log.Error().Msg(err.Error())
	}

	return res, err
}
