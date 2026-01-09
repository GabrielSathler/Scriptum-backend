package routes

import (
	"github.com/GabrielSathler/articles-backend/internal/controller"
	"github.com/GabrielSathler/articles-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup, userService service.UserService) {
	userController := controller.NewUserController(userService)
	deleteController := controller.NewDeleteUserController(userService)

	r.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	r.GET("getUserById/:userId", controller.FindUserById)
	r.GET("getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("createUser", userController.CreateUser)
	r.PUT("updateUser/:userId", controller.UpdateUser)
	r.DELETE("deleteUser/:userId", deleteController.DeleteUser)
}
