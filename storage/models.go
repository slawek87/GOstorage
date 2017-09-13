package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/slawek87/GOstorage/service"
	"github.com/slawek87/GOstorage/settings"
	"os"
	"strconv"
	"mime/multipart"
	"io"
)

// Structure contains all method needed to save / delete / update files.
type FileManager struct {}

// Method checks if current file exist in storage. If it's true then returns true.
func (fileManager *FileManager) IsFileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Method returns full path to stored file.
func (fileManager *FileManager) GetPath(filename string, token string) string {
	return settings.FILE_STORAGE + token + "/" + filename
}

// Method returns available name for filename.
// If basic filename is not available method creates new unique name for that file.
func (fileManager *FileManager) GetFilename(filename string, token string, counter int) string {
	strCounter := counter

	if counter > 0 {
		filename = strconv.Itoa(strCounter) + "_" + filename
	}

	path := fileManager.GetPath(filename, token)

	if fileManager.IsFileExists(path) == true {
		counter++
		return fileManager.GetFilename(filename, token, counter)
	}

	return filename
}

// Method saves file in storage.
func (storage *Storage) SaveFile(file multipart.File, header *multipart.FileHeader, token string) (string, error) {
	filename := storage.FileManager.GetFilename(header.Filename, token, 0)
	path := storage.FileManager.GetPath(filename, token)

	out, err := os.Create(path)

	if err != nil {
		return header.Filename, err
	}
	defer out.Close()

	_, err = io.Copy(out, file)

	return filename, err
}


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