// internal/repository/user_repository.go
package repository

import (
	"github.com/GabrielSathler/articles-backend/internal/controller/model/response"
	"gorm.io/gorm"
)

// UserRepository define os métodos de acesso ao banco
type UserRepository interface {
	Create(user *entity.UserResponse) error
	FindByEmail(email string) (*entity.UserResponse, error)
	FindById(id uint) (*entity.UserResponse, error)
	DeleteUser(id uint) (*entity.UserResponse, error)
}

// userRepositoryImpl implementa a interface
type userRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância do repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

// Create insere um novo usuário no banco
func (r *userRepositoryImpl) Create(user *entity.UserResponse) error {
	return r.db.Create(user).Error
}

// FindByEmail busca usuário por email
func (r *userRepositoryImpl) FindByEmail(email string) (*entity.UserResponse, error) {
	var user entity.UserResponse
	err := r.db.Where("email = ?", email).First(&user).Error
	
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Retorna nil se não encontrar
		}
		return nil, err
	}
	
	return &user, nil
}

// FindById busca usuário por ID
func (r *userRepositoryImpl) FindById(id uint) (*entity.UserResponse, error) {
	var user entity.UserResponse
	err := r.db.First(&user, id).Error
	
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	
	return &user, nil
}

func (r *userRepositoryImpl) DeleteUser(id uint) (*entity.UserResponse, error) {
	var user entity.UserResponse
	err := r.db.Delete(&user, id).Error
	
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}