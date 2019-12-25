package auth

import (
	"context"
	"encoding/json"
	"net/http"
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

		r.Method("POST", "/auth/v1/login", httptransport.NewServer(
			makeLoginEndpoint(svc),
			decodeLoginRequest(log),
			encodeHTTPResponse,
			options...))

		r.Method("POST", "/auth/v1/refresh", httptransport.NewServer(
			makeRefreshEndpoint(svc),
			decodeRefreshTokenRequest(log),
			encodeHTTPResponse,
			options...))
	})

	return r
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

func decodeRefreshTokenRequest(log *zerolog.Logger) httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var req refreshTokenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error().Err(err).Msg("failed to decode request")
			return nil, err
		}

		if err := validator.New().Struct(&req); err != nil {
			log.Error().Err(err).Msg("error validating request")
			return nil, err
		}

		return &req, nil
	}
}

func decodeLoginRequest(log *zerolog.Logger) httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var req loginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error().Err(err).Msg("failed to decode request")
			return nil, err
		}

		//ip, _, err := net.SplitHostPort(r.RemoteAddr)
		//if err != nil {
		//	log.Error().Err(err).Msg("failed to parse host")
		//	return nil, fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
		//}
		//
		//userIP := net.ParseIP(ip)
		//if userIP == nil {
		//	log.Error().Err(err).Msg("failed to parse ip")
		//	return nil, fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
		//}
		//
		//req.IP = userIP
		//req.Ua = r.Header.Get("User-Agent")

		if err := validator.New().Struct(&req); err != nil {
			log.Error().Err(err).Msg("error validating request")
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
