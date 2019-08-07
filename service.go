package backend

import (
	"gopkg.in/go-playground/validator.v9"

	"github.com/dynastiateam/backend/repository"
)

type Service interface {
	AddRequest(req *repository.Request) (*repository.Request, error)
}

type service struct {
	repo repository.Repository
}

const (
	EventCreated   = iota //new event
	EventAccepted         //accepted by guard
	EventCompleted        //competed by guard
	EventRejected         //rejected by guard
	EventMissed           //wasn't completed before ETA
)

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) AddRequest(req *repository.Request) (*repository.Request, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	req.Status = EventCreated

	return s.repo.AddRequest(req)
}
