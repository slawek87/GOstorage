package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slawek87/GOstorage/storage"
	"github.com/slawek87/GOstorage/service"
	"github.com/slawek87/GOstorage/auth"
)


func main() {
	service.InitMigrations()
	storage.InitMigrations()

	goauth := auth.GOauth()
	r := gin.Default()

	v1 := r.Group("api/v1/")
	{
			v1.POST("/service/register", service.ServiceRegisterAPI)
			v1.POST("/storage/file/upload",  goauth.AuthenticationMiddleware, storage.UploadFileAPI)
	}

	r.Run(":8070")
}



