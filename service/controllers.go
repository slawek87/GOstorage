package service

import (
	"time"
	"github.com/slawek87/GOstorage/settings"
	"os"
)

func(service *Service) RegisterService(password string) (*Service, error) {
	goauth := settings.GOauth()

	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()
	service.Token, _ = service.GenerateHash()

	db, _ := settings.InitDB()
	db.NewRecord(&service)
	query := db.Create(&service)

	//register user in GOauth API.
	if query.Error == nil {
		goauth.RegisterUser(service.Name, password)

		storageDir := settings.FILE_STORAGE + service.Token
		err := os.Mkdir(storageDir, 755)

		if err != nil {
			return service, err
		}
	}

	return service, query.Error
}
