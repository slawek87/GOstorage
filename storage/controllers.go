package storage

import (
	"time"
	"github.com/slawek87/GOstorage/settings"
)

func(service *Service) RegisterService(password string) (*Service, error) {
	goauth := settings.GOauth()

	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()
	service.Storage, _ = service.GenerateHash()

	db, _ := settings.InitDB()
	db.NewRecord(&service)
	query := db.Create(&service)

	//register user in GOauth API.
	if query.Error == nil {
		goauth.RegisterUser(service.Name, password)
	}

	return service, query.Error
}
