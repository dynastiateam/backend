package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/dynastiateam/backend/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
}

func (r *repository) CreateUser(user *models.User) (*models.User, error) {
	var u models.User
	if err := r.db.Where("email = ?", user.Email).First(&u).Error; !gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user with this email already exists")
	}

	user.SetPassword(hashAndSalt(user.RawPassword))
	user.RawPassword = ""

	if err := r.db.Create(user).Error; err != nil {
		return nil, errors.Wrap(err, "error creating user")
	}

	return user, nil
}

func hashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
