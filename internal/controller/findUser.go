package controller

import (
	"net/http"

	resterr "github.com/GabrielSathler/articles-backend/internal/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {

	userId := c.Param("user_id")



	err := resterr.NewBadRequestError("invalid json body")
	c.JSON(int(err.Status), err)


	c.JSON(http.StatusOK, userId)
}

func FindUserByEmail(c *gin.Context) {
	err := resterr.NewBadRequestError("invalid json body")
	c.JSON(int(err.Status), err)
}
