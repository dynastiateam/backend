package user

import (
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

func (s *service) Create(user *models.User) (*models.User, error) {
	return s.repo.CreateUser(user)
}
