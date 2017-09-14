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

	filename, err := storage.FileManager.SaveFile(file, header, record.Token)

	if err != nil {
		return storage, err
	}

    storage.FileName = filename
    storage.ServiceID = record.ID

	query := db.Save(&storage)

    return storage, query.Error
}

func (storage *Storage) DeleteFile(token string) error {
	var record service.Service
	name, _, _ := DecodeToken(token)

	db, _ := settings.InitDB()
	query := db.Where(&service.Service{Name:name}).First(&record)

	if query.Error != nil {
		return query.Error
	}

	err := storage.FileManager.DeleteFile(storage.FileName, record.Token)

	if err != nil {
		return err
	}

	db.Model(record).Related(&storage)

	query = db.Delete(&storage)

	return query.Error
}