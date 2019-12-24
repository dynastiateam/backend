package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

type Router struct {
	router    chi.Router
	log       *zerolog.Logger
	jwtSecret string
	httpPort  string
}

func New(jwtSecret, port string, log *zerolog.Logger, userSvc, authSvc http.Handler) *Router {
	r := &Router{
		router:    chi.NewRouter(),
		jwtSecret: jwtSecret,
		log:       log,
		httpPort:  port,
	}

	r.router.Route("/v1", func(v1 chi.Router) {
		v1.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.Header().Set("Content-Type", "application/json")
				if r.Method == http.MethodOptions {
					w.WriteHeader(http.StatusOK)
					return
				}
				next.ServeHTTP(w, r)
			})
		})
		v1.Mount("/user", userSvc)
		v1.Mount("/auth", authSvc)
	})

	r.router.Get("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("works"))
	})

	return r
}

func (r *Router) ListenAndServe() error {
	if len(r.httpPort) > 0 {
		r.log.Info().Str("port", r.httpPort).Msg("start http server on port")
		if err := http.ListenAndServe(fmt.Sprintf(":%s", r.httpPort), r.router); err != nil {
			return err
		}
	}
	return nil
}

//if we need requestID
//https://medium.com/@funfoolsuzi/logging-with-request-id-in-go-microservice-21485c6730da
//func (*Router) loggingHandler(next http.Handler) http.Handler {
//	fn := func(w http.ResponseWriter, r *http.Request) {
//		tStart := time.Now()
//		next.ServeHTTP(w, r)
//		tEnd := time.Now()
//		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), tEnd.Sub(tStart))
//	}
//
//	return http.HandlerFunc(fn)
//}

//func (rr *Router) authHandler(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		//List of endpoints that doesn't require auth
//		notAuth := map[string]struct{}{"/v1/test": {}, "/v1/auth/register": {}, "/v1/auth/login": {}}
//		//check if request does not need authentication, serve the request if it doesn't need it
//		if _, ok := notAuth[r.URL.Path]; ok {
//			next.ServeHTTP(w, r)
//			return
//		}
//
//		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header
//
//		if tokenHeader == "" { //AccessToken is missing, returns with error code 403 Unauthorized
//			rr.errorResponse(w, http.StatusBadRequest, errors.New("missing auth token"))
//			return
//		}
//
//		//The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
//		splitted := strings.Split(tokenHeader, " ")
//		if len(splitted) != 2 {
//			rr.errorResponse(w, http.StatusBadRequest, errors.New("invalid/malformed auth token"))
//			return
//		}
//
//		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
//		tk := &models.Token{}
//
//		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
//			return []byte(rr.jwtSecret), nil
//		})
//
//		if err != nil {
//			rr.errorResponse(w, http.StatusBadRequest, errors.New("malformed authentication token"))
//			return
//		}
//
//		if !token.Valid { //AccessToken is invalid, maybe not signed on this server
//			rr.errorResponse(w, http.StatusForbidden, errors.New("token is not valid"))
//			return
//		}
//
//		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
//		ctx := context.WithValue(r.Context(), "user", tk.ID)
//		r = r.WithContext(ctx)
//		next.ServeHTTP(w, r)
//	})
//}
