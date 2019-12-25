package users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeRegisterEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Register(ctx, request.(*userRegisterRequest))
	}
}

func makeUserByEmailAndPasswordEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UserByEmailAndPassword(ctx, request.(*userByEmailAndPasswordRequest))
	}
}

func makeUserByIDRequest(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UserByID(ctx, request.(int))
	}
}
