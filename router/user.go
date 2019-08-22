package router

import (
	"encoding/json"

	"github.com/valyala/fasthttp"

	"github.com/dynastiateam/backend/models"
)

func (r *router) register(ctx *fasthttp.RequestCtx) {
	var u models.User
	if err := json.Unmarshal(ctx.PostBody(), &u); err != nil {
		ctx.Logger().Printf("%s", err)
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	res, err := r.userService.Create(&u)
	if err != nil {
		ctx.Logger().Printf("%s", err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	r.response(ctx, res)
}

func (r *router) login(ctx *fasthttp.RequestCtx) {
	var u models.User
	if err := json.Unmarshal(ctx.PostBody(), &u); err != nil {
		ctx.Logger().Printf("%s", err)
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}
}

//var AuthMiddleware = func(ctx *fasthttp.RequestCtx, next fasthttp.RequestHandler) fasthttp.RequestHandler {
//	return next
//}
//
//var f = func(handler func([]byte) ([]byte, error)) fasthttp.RequestHandler {
//	return func(ctx *fasthttp.RequestCtx) {
//		//data := decode req
//		//res, err := handler(data)
//		//
//		//marshal(res)
//	}
//}
