package user

import (
	"github.com/pkg/errors"

	"github.com/dynastiateam/backend/models"
	"github.com/dynastiateam/backend/repository"
)

type Service interface {
	Create(user *models.User) (*models.User, error)
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetUserByEmailAndPassword(email, password string) (*models.User, error) {
	return nil, nil
}

func (s *service) Create(user *models.User) (*models.User, error) {
	if user.Email == "" || user.RawPassword == "" || user.Apartment == 0 || user.FirstName == "" || user.LastName == "" {
		return nil, errors.New("fields: email, password, apartment, first_name, last_name are mandatory")
	}

	u, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, errors.Wrap(err, "error creating user")
	}

	return u, nil
}
