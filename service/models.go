package service

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Service struct {
	gorm.Model

	Name  		string 		`gorm:"not null;unique"`
	Token 		string 		`gorm:"not null;unique"`
}

func (service *Service) GenerateHash() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(service.Name), 14)
	return strings.Replace(string(bytes), "/", "", -1), err
}
