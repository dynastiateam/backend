package users

import (
	"context"
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
	Register(ctx context.Context, req *userRegisterRequest) (*userRegisterResponse, error)
	UserByEmailAndPassword(ctx context.Context, r *userByEmailAndPasswordRequest) (*User, error)
	UserByID(ctx context.Context, id int) (*userByIDResponse, error)
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

type userByEmailAndPasswordRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type userByIDResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type userRegisterResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
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

type service struct {
	db *gorm.DB
}

func NewService(log *zerolog.Logger, db *gorm.DB) Service {
	s := &service{
		db: db,
	}
	svc := newLoggingMiddleware(log, s)

	return svc
}

func (s *service) UserByID(_ context.Context, id int) (*userByIDResponse, error) {
	var u User
	if err := s.db.Where("id = ?", id).First(&u).Error; err != nil {
		return nil, err
	}

	return &userByIDResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone,
		Email:     u.Email,
	}, nil
}

func (s *service) UserByEmailAndPassword(ctx context.Context, r *userByEmailAndPasswordRequest) (*User, error) {
	u, err := s.userByEmail(ctx, r.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(r.Password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		//Password does not match!
		return nil, errInvalidCredentials
	}

	return u, nil
}

func (s *service) userByEmail(_ context.Context, email string) (*User, error) {
	var u User
	if err := s.db.Where("email = ?", email).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errUserNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (s *service) Register(ctx context.Context, r *userRegisterRequest) (*userRegisterResponse, error) {
	u, err := s.userByEmail(ctx, r.Email)
	if err != nil && err != errUserNotFound {
		return nil, err
	}

	if u != nil {
		return nil, errUserEmailExists
	}

	pwd, err := s.hashAndSalt(r.Password)
	if err != nil {
		return nil, err
	}

	usr := User{
		Apartment:  r.Apartment,
		BuildingID: r.BuildingID,
		Email:      r.Email,
		Phone:      r.Phone,
		FirstName:  r.FirstName,
		LastName:   r.LastName,
		Password:   pwd,
		Role:       defaultUserRole,
	}

	if err := s.db.Create(&usr).Error; err != nil {
		return nil, err
	}

	return &userRegisterResponse{
		ID:    usr.ID,
		Email: usr.Email,
	}, nil
}

func (s *service) hashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
