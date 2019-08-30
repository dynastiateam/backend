package user

import (
	"log"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/dynastiateam/backend/models"
	"github.com/dynastiateam/backend/repository"
)

type Service interface {
	Create(user *models.User) (*models.User, error)
	Login(email, password string) (*models.User, error)
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Login(email, password string) (*models.User, error) {
	user, err := s.repo.UserByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "error on login user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return nil, errors.New("Invalid login credentials. Please try again")
	}
	user.Password = ""

	return user, nil
}

func (s *service) Create(user *models.User) (*models.User, error) {
	user.Password = hashAndSalt(user.RawPassword)
	user.RawPassword = ""

	u, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, errors.Wrap(err, "error creating user")
	}

	return u, nil
}

func hashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
