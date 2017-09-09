package storage

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	gorm.Model
	Name    	string `gorm:"not null;unique"`
	Storage 	string `gorm:"not null;unique"`
}


func (service *Service) GenerateHash() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(service.Name), 14)
	return string(bytes), err
}
