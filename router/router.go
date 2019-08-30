package router

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	"github.com/dynastiateam/backend/models"
	"github.com/dynastiateam/backend/services/user"

	"github.com/valyala/fasthttp"
)

type Router interface {
	Handler(ctx *fasthttp.RequestCtx)
}

type router struct {
	*mux.Router
	userService user.Service
}

func New(userSvc user.Service) http.Handler {
	r := router{mux.NewRouter(), userSvc}

	r.Methods("POST").Path("/api/v1/auth/register").HandlerFunc(r.register)
	r.Methods("POST").Path("/api/v1/auth/login").HandlerFunc(r.login)
	r.Methods("POST").Path("/api/v1/test").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ahha"))
	})

	r.Use(r.jwtAuth)

	return r
}

func (rr *router) response(w http.ResponseWriter, result interface{}) {
	resp, err := json.Marshal(result)
	if err != nil {
		rr.errorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(resp)
}

func (*router) errorResponse(w http.ResponseWriter, err error) {
	log.Println(err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func (rr *router) jwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//List of endpoints that doesn't require auth
		notAuth := map[string]struct{}{"/api/v1/auth/register": {}, "/api/v1/auth/login": {}}
		//check if request does not need authentication, serve the request if it doesn't need it
		if _, ok := notAuth[r.URL.Path]; ok {
			next.ServeHTTP(w, r)
			return
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			rr.errorResponse(w, errors.New("missing auth token"))
			return
		}

		//The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			rr.errorResponse(w, errors.New("invalid/malformed auth token"))
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			rr.errorResponse(w, errors.New("malformed authentication token"))
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			rr.errorResponse(w, errors.New("token is not valid"))
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		ctx := context.WithValue(r.Context(), "user", tk.ID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
