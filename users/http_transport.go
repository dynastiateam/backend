package users

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rs/zerolog"
	"gopkg.in/go-playground/validator.v9"
)

func NewHTTPHandler(svc Service, log *zerolog.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeHTTPError),
	}

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(accessLogMiddleware(log))

		r.Method("POST", "/users/v1/register", httptransport.NewServer(
			makeRegisterEndpoint(svc),
			decodeRegisterRequest(log),
			encodeHTTPResponse,
			options...))

		r.Method("GET", "/users/v1/user", httptransport.NewServer(
			makeUserByEmailAndPasswordEndpoint(svc),
			decodeUserByEmailAndPasswordRequest(log),
			encodeHTTPResponse,
			options...))

		r.Method("GET", "/users/v1/user/{id}", httptransport.NewServer(
			makeUserByIDRequest(svc),
			decodeUserByIDRequest,
			encodeHTTPResponse,
			options...))
	})

	return r
}

func decodeUserByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" || idStr == "0" {
		return "", errors.New("empty id")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return "", err
	}

	return id, nil
}

func decodeUserByEmailAndPasswordRequest(log *zerolog.Logger) httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (request interface{}, err error) {
		req := userByEmailAndPasswordRequest{
			Email:    r.URL.Query().Get("email"),
			Password: r.URL.Query().Get("pwd"),
		}
		if err := validator.New().Struct(&req); err != nil {
			log.Error().Msg(err.Error())
			return nil, err
		}

		return &req, nil
	}
}

func decodeRegisterRequest(log *zerolog.Logger) httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var req userRegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error().Err(err).Msg("failed to decode request")
			return nil, errors.New("failed to decode request")
		}

		if err := validator.New().Struct(&req); err != nil {
			return nil, err
		}
		return &req, nil
	}

}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(response)
}

func encodeHTTPError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

var accessLogMiddleware = func(log *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r)
			duration := time.Since(start)
			log.Info().
				Str("tag", "http_log").
				Str("remote_addr", r.RemoteAddr).
				Str("request_method", r.Method).
				Str("request_uri", r.RequestURI).
				Int("response_code", ww.Status()).
				Dur("duration", duration).
				Msg("request")
		})
	}
}
