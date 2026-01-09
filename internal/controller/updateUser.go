package controller

import (
	resterr "github.com/GabrielSathler/articles-backend/internal/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	err := resterr.NewBadRequestError("invalid json body")
	c.JSON(int(err.Status), err)
}
