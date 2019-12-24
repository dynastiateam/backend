package router

//
//import (
//	"encoding/json"
//	"net/http"
//	"os"
//	"time"
//
//	"github.com/dgrijalva/jwt-go"
//	"gopkg.in/go-playground/validator.v9"
//
//	"github.com/dynastiateam/backend/models"
//)
//
//type userRegisterRequest struct {
//	Apartment  int    `json:"apartment,omitempty" validate:"required"`
//	Email      string `json:"email,omitempty" validate:"required,email"`
//	Password   string `json:"password,omitempty" validate:"required"`
//	Phone      string `json:"phone,omitempty" validate:"required"`
//	FirstName  string `json:"first_name,omitempty" validate:"required"`
//	LastName   string `json:"last_name,omitempty" validate:"required"`
//	BuildingID int    `json:"building_id,omitempty" validate:"required"`
//}

//type loginRequest struct {
//	Email    string `json:"email" validate:"required,email"`
//	Password string `json:"password" validate:"required"`
//}
//
//type loginResponse struct {
//	AccessToken  string `json:"access_token"`
//	RefreshToken string `json:"refresh_token"`
//	ExpiresIn    int64  `json:"expires_in"`
//}
//
//func (rr *Router) register(w http.ResponseWriter, r *http.Request) {
//	var req userRegisterRequest
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		rr.errorResponse(w, http.StatusBadRequest, err)
//		return
//	}
//
//	if err := validator.New().Struct(&req); err != nil {
//		rr.errorResponse(w, http.StatusBadRequest, err)
//		return
//	}
//
//	u := models.User{
//		Apartment:   req.Apartment,
//		Email:       req.Email,
//		Phone:       req.Phone,
//		FirstName:   req.FirstName,
//		LastName:    req.LastName,
//		BuildingID:  req.BuildingID,
//		RawPassword: req.Password,
//		Role:        4,
//	}
//
//	res, err := rr.userService.Create(&u)
//	if err != nil {
//		rr.errorResponse(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	rr.response(w, res)
//}
//
//func (rr *Router) login(w http.ResponseWriter, r *http.Request) {
//	var req loginRequest
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		rr.errorResponse(w, http.StatusBadRequest, err)
//		return
//	}
//
//	if err := validator.New().Struct(&req); err != nil {
//		rr.errorResponse(w, http.StatusBadRequest, err)
//		return
//	}
//
//	user, err := rr.userService.Login(req.Email, req.Password)
//	if err != nil {
//		rr.errorResponse(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	refreshToken := &models.Token{
//		ID: user.ID,
//		StandardClaims: jwt.StandardClaims{
//			Audience:  "dynapp",
//			ExpiresAt: time.Now().Add(60 * 24 * time.Hour).Unix(), //60 days
//			Issuer:    "auth.dynapp",
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), refreshToken)
//	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

//	res := loginResponse{
//		AccessToken:  "",
//		RefreshToken: tokenString,
//		ExpiresIn:    0,
//	}
//
//	rr.response(w, res)
//}
