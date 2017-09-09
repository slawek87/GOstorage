package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slawek87/GOstorage/storage"
)


func main() {
	storage.InitMigrations()

	r := gin.Default()

	v1 := r.Group("api/v1/storage/")
	{
			v1.POST("/service/register", storage.ServiceRegisterAPI)
	}

	r.Run(":8070")
}



