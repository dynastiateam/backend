package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"

	"github.com/dynastiateam/backend/services/user"
)

var (
	errLoginFailed = errors.New("failed to login user")
)

type Service interface {
	Login(request loginRequest) (*loginResponse, error)
}

type UserService interface {
	UserByEmailAndPassword(email, password string) (*user.User, error)
}

type service struct {
	log       *zerolog.Logger
	uSrv      UserService
	repo      Repository
	jwtSecret string
}

type session struct {
	ID           string
	UserID       int
	RefreshToken uuid.UUID
	UA           string
	Fingerprint  string
	IP           string
	ExpiresIn    int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewService(log *zerolog.Logger, db *gorm.DB, uSrv UserService, jwtSecret string) Service {
	return &service{
		log:       log,
		uSrv:      uSrv,
		repo:      newRepo(db),
		jwtSecret: jwtSecret,
	}
}

func (s *service) Login(r loginRequest) (*loginResponse, error) {
	u, err := s.uSrv.UserByEmailAndPassword(r.Email, r.Password)
	if err != nil {
		return nil, err
	}

	sess, err := s.repo.CreateSession(&r, u.ID)
	if err != nil {
		s.log.Error().Err(err)
		return nil, errLoginFailed
	}

	t := &Token{
		ID:   u.ID,
		Name: fmt.Sprintf("%s %s", u.FirstName, u.LastName),
		Role: u.Role,
		StandardClaims: jwt.StandardClaims{
			Audience:  "dynapp",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "auth.dynapp",
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenString, _ := token.SignedString([]byte(s.jwtSecret))

	return &loginResponse{
		AccessToken:  tokenString,
		RefreshToken: sess.RefreshToken.String(),
	}, nil
}
