package controller

import (
	"github.com/gin-gonic/gin"
	"myBulebell/pkg/serializer"
	"myBulebell/service/user"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var service user.RegisterService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Err(serializer.CodeParamErr, err.Error(), err))
	}
}

func UserActive(c *gin.Context) {
	var service user.ActiveService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Active(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Err(serializer.CodeParamErr, err.Error(), err))
	}
}
