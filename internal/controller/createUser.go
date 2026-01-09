// internal/controller/createUser.go
package controller

import (
	"net/http"

	resterr "github.com/GabrielSathler/articles-backend/internal/configuration/rest_err"
	models "github.com/GabrielSathler/articles-backend/internal/controller/model/request"
	"github.com/GabrielSathler/articles-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UserController encapsula as dependências do controller
type UserController struct {
	userService service.UserService
}

// NewUserController cria uma nova instância do controller
func NewUserController(service service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// CreateUser é o handler HTTP para criar usuário
func (ctrl *UserController) CreateUser(c *gin.Context) {
	
	// 1. Fazer o binding do JSON para a struct UserRequest
	var userRequest models.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		
		// Verificar se é erro de validação
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			causes := []resterr.Causes{}
			
			for _, e := range validationErrors {
				cause := resterr.Causes{
					Field:   e.Field(),
					Message: getValidationMessage(e),
				}
				causes = append(causes, cause)
			}
			
			restErr := resterr.NewBadRequestValidationError("invalid fields", causes)
			c.JSON(int(restErr.Status), restErr)
			return
		}
		
		// Erro genérico de JSON
		restErr := resterr.NewBadRequestError("invalid json body")
		c.JSON(int(restErr.Status), restErr)
		return
	}
	
	// 2. Chamar o service para criar o usuário
	userResponse, restErr := ctrl.userService.CreateUser(userRequest)
	if restErr != nil {
		c.JSON(int(restErr.Status), restErr)
		return
	}
	
	// 3. Retornar sucesso com status 201 Created
	c.JSON(http.StatusCreated, userResponse)
}

// getValidationMessage retorna mensagens amigáveis para erros de validação
func getValidationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "invalid email format"
	case "min":
		return "minimum length is " + e.Param()
	case "max":
		return "maximum length is " + e.Param()
	case "oneof":
		return "must be one of: " + e.Param()
	default:
		return "invalid value"
	}
}