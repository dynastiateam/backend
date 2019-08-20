package router

import (
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

type RegisterRequest struct {
	Phone     string `json:"phone" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func (r *RegisterRequest) Validate() error {
	return validator.New().Struct(r)
}

func Register(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {

	})
	//var req RegisterRequest
	//if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
	//	ctx.Logger().Printf("%s", err)
	//	ctx.Error(err.Error(), fasthttp.StatusBadRequest)
	//	return
	//}
	//
	//if err := validator.New().Struct(req); err != nil {
	//	ctx.Logger().Printf("%s", err)
	//	ctx.Error(err.Error(), fasthttp.StatusBadRequest)
	//	return
	//}

}
