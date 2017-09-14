package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/slawek87/GOstorage/service"
	"github.com/slawek87/GOstorage/settings"
)

// Model for Storage
type Storage struct {
	gorm.Model

	FileName    string  `gorm:"not null"`

	Service     service.Service
	ServiceID 	uint	`gorm:"not null"`

	FileManager FileManager
}

// Method returns full path to http storage.
func (storage *Storage) GetUrl() string {
	getService := service.Service{}

	db, _ := settings.InitDB()
	db.Model(storage).Related(&getService)

	return "/storage/" + getService.Token + "/" + storage.FileName
}