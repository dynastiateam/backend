package auth

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	router chi.Router
	log    *zerolog.Logger
	srv    Service
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	IP       net.IP `json:"ip" validate:"required"`
	Ua       string `json:"ua" validate:"required"`
}

type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Token struct {
	ID   int
	Name string
	Role int
	jwt.StandardClaims
}

func NewHandler(log *zerolog.Logger, svc Service) *Handler {
	r := &Handler{log: log, router: chi.NewRouter(), srv: svc}

	return r
}

func (h *Handler) Routes() http.Handler {
	h.router.Post("/login", h.login)

	return h.router
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error().Err(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to unmarshal request"))
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		h.log.Error().Err(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("userip: %q is not IP:port", r.RemoteAddr)))
		return
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		h.log.Error().Err(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("userip: %q is not IP:port", r.RemoteAddr)))
		return
	}

	req.IP = userIP
	req.Ua = r.Header.Get("User-Agent")

	if err := validator.New().Struct(&req); err != nil {
		h.log.Error().Err(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := h.srv.Login(req)
	if err != nil {
		h.log.Error().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to login user"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": res,
	})
}
