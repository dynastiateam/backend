package router

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/go-playground/validator.v9"

	"github.com/dynastiateam/backend/models"
)

type userRegisterRequest struct {
	Apartment  int    `json:"apartment,omitempty" validate:"required"`
	Email      string `json:"email,omitempty" validate:"required,email"`
	Password   string `json:"password,omitempty" validate:"required"`
	Phone      string `json:"phone,omitempty" validate:"required"`
	FirstName  string `json:"first_name,omitempty" validate:"required"`
	LastName   string `json:"last_name,omitempty" validate:"required"`
	BuildingID int    `json:"building_id,omitempty" validate:"required"`
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (rr *router) register(w http.ResponseWriter, r *http.Request) {
	var req userRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rr.errorResponse(w, err)
		return
	}

	if err := validator.New().Struct(&req); err != nil {
		rr.errorResponse(w, err)
		return
	}

	u := models.User{
		Apartment:   req.Apartment,
		Email:       req.Email,
		Phone:       req.Phone,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		BuildingID:  req.BuildingID,
		RawPassword: req.Password,
	}

	res, err := rr.userService.Create(&u)
	if err != nil {
		rr.errorResponse(w, err)
		return
	}

	rr.response(w, res)
}

func (rr *router) login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rr.errorResponse(w, err)
		return
	}

	if err := validator.New().Struct(&req); err != nil {
		rr.errorResponse(w, err)
		return
	}

	res, err := rr.userService.Login(req.Email, req.Password)
	if err != nil {
		rr.errorResponse(w, err)
		return
	}

	tk := &models.Token{ID: res.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	res.Token = tokenString

	rr.response(w, res)
}

//var AuthMiddleware = func(ctx *fasthttp.RequestCtx, next fasthttp.RequestHandler) fasthttp.RequestHandler {
//	return next
//}
//
//var f = func(handler func([]byte) ([]byte, error)) fasthttp.RequestHandler {
//	return func(ctx *fasthttp.RequestCtx) {
//		//data := decode req
//		//res, err := handler(data)
//		//
//		//marshal(res)
//	}
//}
