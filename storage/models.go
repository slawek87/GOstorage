package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/slawek87/GOstorage/service"
	"github.com/slawek87/GOstorage/settings"
)


type Storage struct {
	gorm.Model

	FileName    string  `gorm:"not null"`

	Service     service.Service
	ServiceID 	uint	`gorm:"not null"`

}


func (storage *Storage) GetUrl() string {
	getService := service.Service{}

	db, _ := settings.InitDB()
	db.Model(storage).Related(&getService)

	return "/storage/" + getService.Token + "/" + storage.FileName
}