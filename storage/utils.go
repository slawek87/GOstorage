package storage

import (
	"strings"
	"errors"
	"encoding/base64"
)

// method decodes Authorisation token. Each token include username and password.
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