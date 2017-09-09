package storage

import "github.com/slawek87/GOstorage/settings"


func InitMigrations() {
	db, _ := settings.InitDB()

	db.LogMode(true)
	db.AutoMigrate(&Service{})
}

