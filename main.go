package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slawek87/GOstorage/storage"
	"github.com/slawek87/GOstorage/settings"
)


func main() {
	storage.InitMigrations()
	goauth := settings.GOauth()

	r := gin.Default()

	v1 := r.Group("api/v1/storage/")
	{
		v1.POST("/service/register/", goauth.AuthenticationMiddleware, storage.ServiceRegisterAPI)
	}

	r.Run(":8070")
}



