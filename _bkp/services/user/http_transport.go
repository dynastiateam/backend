package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

type Handler struct {
	router chi.Router
	log    *zerolog.Logger
	srv    Service
}

func NewHandler(log *zerolog.Logger, svc Service) *Handler {
	r := &Handler{log: log, router: chi.NewRouter(), srv: svc}

	return r
}

func (h *Handler) Routes() http.Handler {
	h.router.Post("/register", h.register)

	return h.router
}

type userRegisterRequest struct {
	Apartment  int    `json:"apartment,omitempty" validate:"required"`
	Email      string `json:"email,omitempty" validate:"required,email"`
	Password   string `json:"password,omitempty" validate:"required"`
	Phone      string `json:"phone,omitempty" validate:"required"`
	FirstName  string `json:"first_name,omitempty" validate:"required"`
	LastName   string `json:"last_name,omitempty" validate:"required"`
	BuildingID int    `json:"building_id,omitempty" validate:"required"`
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var req userRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error().Err(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to unmarshal request"))
		return
	}

	if err := validator.New().Struct(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := h.srv.Create(req)
	if err != nil {
		h.log.Error().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create User"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res) //nolint: errcheck
}
