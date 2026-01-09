package database

import (
	"github.com/GabrielSathler/articles-backend/internal/controller/model/response"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.UserResponse) error
	FindByEmail(email string) (*entity.UserResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *entity.UserResponse) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*entity.UserResponse, error) {
	var user entity.UserResponse
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
