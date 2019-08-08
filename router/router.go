package router

import (
	"encoding/json"
	"github.com/dynastiateam/backend/repository"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/dynastiateam/backend"
)

//var mid = func(ctx *fasthttp.RequestCtx, next fasthttp.RequestHandler) fasthttp.RequestHandler {
//	return next
//}

func New(svc backend.Service) *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.POST("/signin", func(ctx *fasthttp.RequestCtx) {
		//https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
		//https://4gophers.ru/articles/avtorizaciya-v-go-s-ispolzovaniem-jwt/#.XUmtZCVn2Ec
		creds := make(map[string]string)
		if err := json.Unmarshal(ctx.PostBody(), &creds); err != nil {
			log.Println(err)
			ctx.Error(err.Error(), fasthttp.StatusBadRequest)
			return
		}

		if _, ok := creds["login"]; !ok {
			log.Println("no login")
			ctx.Error("no login", fasthttp.StatusBadRequest)
			return
		}

		if _, ok := creds["password"]; !ok {
			log.Println("no password")
			ctx.Error("no password", fasthttp.StatusBadRequest)
			return
		}
	})

	router.POST("/request", func(ctx *fasthttp.RequestCtx) {
		var req repository.Request
		if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
			log.Println(err)
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		res, err := svc.AddRequest(&req)
		if err != nil {
			log.Println(err)
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		ctx.SetContentType("application/json; charset=utf8")
		ctx.Response.SetBody(resp)
	})

	return router
}
