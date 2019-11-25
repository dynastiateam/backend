package user

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

const defaultUserRole = 4

var (
	errUserNotFound       = errors.New("user not found")
	errUserEmailExists    = errors.New("user with this email already exists")
	errInvalidCredentials = errors.New("invalid login credentials")
)

type Service interface {
	Create(request userRegisterRequest) (int, error)
	UserByEmailAndPassword(email, password string) (*User, error)
}

type service struct {
	log  *zerolog.Logger
	repo Repository
}

type User struct {
	ID          int    `json:"id"`
	Apartment   int    `json:"apartment,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	Phone       string `json:"phone,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Role        int    `json:"role,omitempty"`
	ResidenceID int    `json:"residence_id,omitempty"`
	BuildingID  int    `json:"building_id,omitempty"`
}

func NewService(log *zerolog.Logger, db *gorm.DB) Service {
	return &service{
		repo: newRepo(db),
		log:  log,
	}
}

func (s *service) Create(r userRegisterRequest) (int, error) {
	u := User{
		Apartment:  r.Apartment,
		BuildingID: r.BuildingID,
		Email:      r.Email,
		Phone:      r.Phone,
		FirstName:  r.FirstName,
		LastName:   r.LastName,
		Role:       defaultUserRole,
	}
	u.Password = s.hashAndSalt(r.Password)

	err := s.repo.CreateUser(&u)
	if err != nil {
		s.log.Error().Err(err)
	}

	return u.ID, err
}

func (s *service) UserByEmailAndPassword(email, password string) (*User, error) {
	u, err := s.repo.UserByEmail(email)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return nil, errInvalidCredentials
	}

	return u, nil
}

func (s *service) hashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		s.log.Error().Err(err)
	}
	return string(hash)
}
