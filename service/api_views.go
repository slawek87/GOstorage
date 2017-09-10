package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServiceRegisterAPI(c *gin.Context) {
	service := Service{}
	c.Bind(&service)

	_, err := service.RegisterService(c.PostForm("Password"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, &service)
}
