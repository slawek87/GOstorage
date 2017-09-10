package storage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFileAPI(c *gin.Context) {
	file, header , err := c.Request.FormFile("upload")
	token := c.GetHeader("Authorization")

	storage := Storage{}
	c.Bind(&storage)

	_, err = storage.UploadFile(file, header, token)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"Uploaded": &storage.FileName})
}

