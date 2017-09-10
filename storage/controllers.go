package storage

import (
	"mime/multipart"
	"errors"
	"github.com/slawek87/GOstorage/settings"
	"github.com/slawek87/GOstorage/service"
)

func (storage *Storage) UploadFile(file multipart.File, header *multipart.FileHeader, token string) (*Storage, error) {
	if file == nil {
		return storage, errors.New("No file.")
	}

	var record service.Service
	name, _, _ := DecodeToken(token)

	db, _ := settings.InitDB()
	db.Where(&service.Service{Name:name}).First(&record)

	filename, err := storage.SaveFile(file, header, record.Token)

	if err != nil {
		return storage, err
	}

    storage.FileName = filename
    storage.ServiceID = record.ID

	query := db.Save(&storage)

    return storage, query.Error
}