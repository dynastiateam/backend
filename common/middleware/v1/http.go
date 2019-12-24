package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog"

	"github.com/go-chi/chi/middleware"
)

const (
	//RequestHeaderName header name for correlation id
	RequestHeaderName = "X-Amzn-Trace-Id"
	//RequestContextKey key name for context where correlation id stored
	RequestContextKey = "request_id"
	//RequestURIName
	RequestURI = "RequestURI"
)

func RequestIDFromHeaderToCTX(ctx context.Context, r *http.Request) context.Context {
	if v := r.Header.Get(RequestHeaderName); v != "" {
		ctx = context.WithValue(ctx, RequestContextKey, v)
	}

	return ctx
}

func RequestURIToContext(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, RequestURI, r.URL)
}

// middlewareLog holds dependencies for request information logging.
type middlewareLog struct {
	log *zerolog.Logger
}

// NewMiddlewareLog return middleware.
func NewMiddlewareLog(log *zerolog.Logger) func(next http.Handler) http.Handler {
	mid := &middlewareLog{
		log: log,
	}
	return mid.middleware
}

// middleware represent logging middleware logic.
func (m *middlewareLog) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		duration := time.Since(start)
		m.log.Info().
			Str("tag", "http_log").
			Str("remote_addr", r.RemoteAddr).
			Str("request_method", r.Method).
			Str("request_uri", r.RequestURI).
			Int("response_code", ww.Status()).
			Dur("duration", duration).
			Msg("request")
	}

	return http.HandlerFunc(fn)
}
