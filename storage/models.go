package storage

import (
	"github.com/jinzhu/gorm"
)


type Storage struct {
	gorm.Model

	FileName    string  `gorm:"not null"`
	ServiceID 	uint	`gorm:"null;unique"`
}
