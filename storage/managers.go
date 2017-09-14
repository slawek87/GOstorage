package storage

import (
	"os"
	"strconv"
	"mime/multipart"
	"io"
	"github.com/slawek87/GOstorage/settings"
)

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
func (fileManager *FileManager) SaveFile(file multipart.File, header *multipart.FileHeader, token string) (string, error) {
	filename := fileManager.GetFilename(header.Filename, token, 0)
	path := fileManager.GetPath(filename, token)

	out, err := os.Create(path)

	if err != nil {
		return header.Filename, err
	}
	defer out.Close()

	_, err = io.Copy(out, file)

	return filename, err
}

// method delete file from storage - file system and db.
func (fileManager *FileManager) DeleteFile(filename string, token string) error {
	path := fileManager.GetPath(filename, token)

	err := os.Remove(path)

	return err
}




