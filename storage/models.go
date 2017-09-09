package storage

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model
	Name    	string
	Storage 	string
}
