package service

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Service struct {
	gorm.Model

	Name  		string 		`gorm:"not null;unique"`
	Token 		string 		`gorm:"not null;unique"`
}

func (service *Service) GenerateToken() string {
	return uuid.NewV4().String()
}
