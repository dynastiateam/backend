package router

import (
	"encoding/json"

	"github.com/dynastiateam/backend/models"
	"github.com/dynastiateam/backend/services/user"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type Router interface {
	Handler(ctx *fasthttp.RequestCtx)
}

type router struct {
	*fasthttprouter.Router
}

//var AuthMiddleware = func(ctx *fasthttp.RequestCtx, next fasthttp.RequestHandler) fasthttp.RequestHandler {
//	return next
//}

func New(userSvc user.Service) Router {
	r := router{fasthttprouter.New()}

	r.POST("/api/v1/auth/register", func(ctx *fasthttp.RequestCtx) {
		var u models.User
		if err := json.Unmarshal(ctx.PostBody(), &u); err != nil {
			ctx.Logger().Printf("%s", err)
			ctx.Error(err.Error(), fasthttp.StatusBadRequest)
			return
		}

		res, err := userSvc.Create(&u)
		if err != nil {
			ctx.Logger().Printf("%s", err)
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		r.response(ctx, res)
	})

	//////////////////////////////////////////

	//router.POST("/signin", func(ctx *fasthttp.RequestCtx) {
	//	//https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
	//	//https://4gophers.ru/articles/avtorizaciya-v-go-s-ispolzovaniem-jwt/#.XUmtZCVn2Ec
	//	creds := make(map[string]string)
	//	if err := json.Unmarshal(ctx.PostBody(), &creds); err != nil {
	//		ctx.Logger().Printf("%s", err)
	//		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
	//		return
	//	}
	//
	//	if _, ok := creds["login"]; !ok {
	//		ctx.Logger().Printf("no login")
	//		ctx.Error("no login", fasthttp.StatusBadRequest)
	//		return
	//	}
	//
	//	if _, ok := creds["password"]; !ok {
	//		ctx.Logger().Printf("no password")
	//		ctx.Error("no password", fasthttp.StatusBadRequest)
	//		return
	//	}
	//})
	//
	//router.POST("/request", func(ctx *fasthttp.RequestCtx) {
	//	var req repository.Request
	//	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
	//		log.Println(err)
	//		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	//		return
	//	}
	//
	//	res, err := svc.AddRequest(&req)
	//	if err != nil {
	//		log.Println(err)
	//		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	//		return
	//	}
	//
	//	resp, err := json.Marshal(res)
	//	if err != nil {
	//		log.Println(err)
	//		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	//		return
	//	}
	//
	//	ctx.SetContentType("application/json; charset=utf8")
	//	ctx.Response.SetBody(resp)
	//})

	return r
}

func (*router) response(ctx *fasthttp.RequestCtx, result interface{}) {
	resp, err := json.Marshal(result)
	if err != nil {
		ctx.Logger().Printf("%s", err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.Success("application/json; charset=utf8", resp)
}
