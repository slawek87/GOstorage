package storage

import (
	"mime/multipart"
	"io"
	"os"
	"github.com/slawek87/GOstorage/settings"
	"strings"
	"errors"
	"encoding/base64"
	"strconv"
)

func (storage *Storage) IsFileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (storage *Storage) GetPath(filename string, token string) string {
	return settings.FILE_STORAGE + token + "/" + filename
}

func (storage *Storage) GetFilename(filename string, token string, counter int) string {
	strCounter := counter

	if counter > 0 {
		filename = strconv.Itoa(strCounter) + "_" + filename
	}

	path := storage.GetPath(filename, token)

	if storage.IsFileExists(path) == true {
		counter++
		return storage.GetFilename(filename, token, counter)
	}

	return filename
}

func (storage *Storage) SaveFile(file multipart.File, header *multipart.FileHeader, token string) (string, error) {
    filename := storage.GetFilename(header.Filename, token, 0)
    path := storage.GetPath(filename, token)

	out, err := os.Create(path)

	if err != nil {
		return header.Filename, err
	}
	defer out.Close()

	_, err = io.Copy(out, file)

	return filename, err
}

func DecodeToken(token string) (string, string, error) {
	splitToken := strings.Split(token, " ")

	if len(splitToken) == 0 {
		return "", "", errors.New("Empty token.")
	}

	decodeToken, err := base64.StdEncoding.DecodeString(splitToken[1])

	if err != nil {
		return "", "", err
	}

	splitDecodeToken := strings.Split(string(decodeToken), ":")

	return splitDecodeToken[0], splitDecodeToken[1], err
}