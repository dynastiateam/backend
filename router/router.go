package router

import (
	"encoding/json"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/dynastiateam/backend"
)

func New(svc backend.Service) *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.POST("/request", func(ctx *fasthttp.RequestCtx) {
		var req backend.Request
		if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
			log.Println(err)
			ctx.Error(err.Error(), fasthttp.StatusNotFound)
			return
		}

		res, err := svc.AddRequest(req)
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
