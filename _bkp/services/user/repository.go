package user

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	CreateUser(u *User) error
	UserByEmail(email string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) UserByEmail(email string) (*User, error) {
	var u User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *repository) CreateUser(user *User) error {
	u, err := r.UserByEmail(user.Email)
	if err != nil {
		return err
	}
	if u != nil {
		return errUserEmailExists
	}

	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
