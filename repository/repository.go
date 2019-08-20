package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	UserRepository
}

type repository struct {
	db *gorm.DB
}

//type Request struct {
//	ID     int `json:"id"`
//	UserID int `json:"user_id" validate:"required"`
//	Type   int `json:"type" validate:"required"`
//	ETA    int `json:"eta" validate:"required"`
//	Status int `json:"status"`
//}

func New(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
