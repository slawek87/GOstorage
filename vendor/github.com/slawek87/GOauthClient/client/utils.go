package client

import (
	"github.com/slawek87/GOauthClient/storage"
	"time"
	"github.com/gin-gonic/gin"
	"strings"
	"errors"
	"encoding/base64"
)

func (goAuth *GOauth) GetURL(url string) string {
	settings := goAuth.Client.Settings
	return settings["PROTOCOL"] + "://" + settings["HOST"] + ":" + settings["PORT"] + url
}

var TokenExpirationTime time.Duration = 8 * time.Hour // 8hrs

func getAuthorizeKey(userID string) string {
	return "goauth" + userID
}

func AuthorizeUser(userID string) {
	redis, _ := storage.RedisDB()

	key := getAuthorizeKey(userID)
	redis.Set(key, true, TokenExpirationTime)
}

func IsAuthorizedUser(userID string) bool {
	redis, _ := storage.RedisDB()

	key := getAuthorizeKey(userID)
	value, _ := redis.Get(key).Result()

	if value != "" {
		return true
	}

	return false
}

func Forbidden(c *gin.Context, err error) *gin.Context {
	c.JSON(401, err)
	c.Abort()

	return c
}

func decodeToken(token string) (string, string, error) {
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


