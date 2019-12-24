package auth

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	CreateSession(req *loginRequest, userID int) (*session, error)
}

type repository struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateSession(req *loginRequest, userID int) (*session, error) {
	//todo do not create session if user already has one
	rt := uuid.NewV4()

	sess := session{
		UserID:       userID,
		RefreshToken: rt,
		UA:           req.Ua,
		IP:           req.IP.String(),
		ExpiresIn:    time.Now().Add(100 * 365 * 24 * time.Hour).Unix(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := r.db.Create(&sess).Error; err != nil {
		return nil, err
	}

	return &sess, nil
}
