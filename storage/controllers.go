package storage

import (
	"time"
	"github.com/slawek87/GOstorage/settings"
)

func(service *Service) RegisterService() (*Service, error) {
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	db, _ := settings.InitDB()
	db.NewRecord(&service)
	query := db.Create(&service)


	return service, query.Error
}
