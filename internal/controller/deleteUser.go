package controller

import (
	"net/http"

	resterr "github.com/GabrielSathler/articles-backend/internal/configuration/rest_err"
	models "github.com/GabrielSathler/articles-backend/internal/controller/model/request"
	"github.com/GabrielSathler/articles-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeleteUserController struct {
	userService service.UserService
}

func NewDeleteUserController(service service.UserService) *DeleteUserController {
	return &DeleteUserController{
		userService: service,
	}
}

func (d *DeleteUserController) DeleteUser(c *gin.Context) {

	var userDelete models.UserRequest
	if err := c.ShouldBindJSON(&userDelete); err != nil {
		if validationErros, ok := err.(validator.ValidationErrors); ok {
			causes := []resterr.Causes{}

			for _, e := range validationErros {
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

		restErr := resterr.NewBadRequestError("invalid json body")
		c.JSON(int(restErr.Status), restErr)
		return
	}

	// chamar o service com o ID e retornar a resposta de `response.UserResponse`
	resp, restErr := d.userService.DeleteUser(userDelete.ID)
	if restErr != nil {
		c.JSON(int(restErr.Status), restErr)
		return
	}

	c.JSON(http.StatusOK, resp)
}
