package repository

import (
	"database/sql"
	"time"
)

type Repository interface {
	AddRequest(req *Request) (*Request, error)
}

type repository struct {
	db *sql.DB
}

type Request struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id" validate:"required"`
	Type   int       `json:"type" validate:"required"`
	ETA    time.Time `json:"eta" validate:"required"`
	Status int       `json:"status"`
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddRequest(req *Request) (*Request, error) {
	return req, nil
}
