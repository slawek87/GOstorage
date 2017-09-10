package storage

import (
	"mime/multipart"
	"io"
	"os"
	"github.com/slawek87/GOstorage/settings"
	"strings"
	"errors"
	"encoding/base64"
)

func (storage *Storage) SaveFile(file multipart.File, header *multipart.FileHeader, token string) (string, error) {
	out, err := os.Create(settings.FILE_STORAGE + token + "/" + header.Filename)

	if err != nil {
		return header.Filename, err
	}
	defer out.Close()

	_, err = io.Copy(out, file)

	return header.Filename, err
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