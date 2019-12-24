package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	middlewares "github.com/dynastiateam/backend/common/middleware/v1"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rs/zerolog"
)

func NewHTTPHandler(svc Service, log *zerolog.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeHTTPError),
		httptransport.ServerBefore(middlewares.RequestIDFromHeaderToCTX),
	}

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middlewares.NewMiddlewareLog(log))

		r.Method("GET", "/user/v1/hello/{key}", httptransport.NewServer(
			makeHelloEndpoint(svc),
			decodeHelloRequest,
			encodeHTTPResponse,
			options...))
	})

	return r
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	key := chi.URLParam(r, "key")

	if key == "" {
		return "", errors.New("empty key")
	}

	return key, nil
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(map[string]interface{}{"message": response})
}

func encodeHTTPError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
}
