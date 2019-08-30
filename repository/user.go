package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/dynastiateam/backend/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	UserByEmail(email string) (*models.User, error)
}

func (r *repository) UserByEmail(email string) (*models.User, error) {
	var u models.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *repository) CreateUser(user *models.User) (*models.User, error) {
	var u models.User
	if err := r.db.Where("email = ?", user.Email).First(&u).Error; !gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user with this email already exists")
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
