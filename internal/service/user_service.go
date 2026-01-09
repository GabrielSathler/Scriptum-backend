// internal/service/user_service.go
package service

import (
	"fmt"
	"time"

	resterr "github.com/GabrielSathler/articles-backend/internal/configuration/rest_err"
	requestModels "github.com/GabrielSathler/articles-backend/internal/controller/model/request"
	responseModels "github.com/GabrielSathler/articles-backend/internal/controller/model/response"
	"github.com/GabrielSathler/articles-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService define os métodos de lógica de negócio
type UserService interface {
	CreateUser(userRequest requestModels.UserRequest) (*responseModels.UserResponse, *resterr.RestErr)
	DeleteUser(userID uint) (responseModels.UserResponse, *resterr.RestErr)
}

// userServiceImpl implementa a interface
type userServiceImpl struct {
	repository repository.UserRepository
}

// NewUserService cria uma nova instância do service
func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{
		repository: repo,
	}
}

// CreateUser contém toda a lógica de criação de usuário
func (s *userServiceImpl) CreateUser(userRequest requestModels.UserRequest) (*responseModels.UserResponse, *resterr.RestErr) {

	// 1. Verificar se o email já existe
	existingUser, err := s.repository.FindByEmail(userRequest.Email)
	if err != nil {
		return nil, resterr.NewInternalServerError("error checking if user exists")
	}

	if existingUser != nil {
		return nil, resterr.NewBadRequestError("email already registered")
	}

	// 2. Fazer hash da senha usando bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userRequest.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, resterr.NewInternalServerError("error encrypting password")
	}

	// 3. Criar a entidade UserResponse para salvar no banco
	// Não atribuir `ID` ou `CreatedAt` aqui: o GORM preenche automaticamente os timestamps e o ID ao salvar.
	user := &responseModels.UserResponse{
		Email:    userRequest.Email,
		Name:     userRequest.Name,
		Password: string(hashedPassword),
		TypeUser: userRequest.TypeUser,
		CreatedAt: time.Now(),
	}

	// 4. Salvar no banco de dados
	if err := s.repository.Create(user); err != nil {
		return nil, resterr.NewInternalServerError(
			fmt.Sprintf("error saving user to database: %s", err.Error()),
		)
	}

	// 5. Retornar a resposta (sem expor o hash da senha)
	// O objeto `user` agora contém `ID`, `CreatedAt` e `UpdatedAt` após o Create.
	response := &responseModels.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		TypeUser:  user.TypeUser,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (s *userServiceImpl) DeleteUser(userID uint) (responseModels.UserResponse, *resterr.RestErr) {
	// 2. Deletar o usuário
	user, err := s.repository.DeleteUser(userID)
	if err != nil {
		return responseModels.UserResponse{}, resterr.NewInternalServerError("error deleting user")
	}

	return *user, nil
}
