package client

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// this is middleware for GIN Framework/
func (goauth *GOauth) AuthenticationMiddleware(c *gin.Context) {
	var err error

	token := c.GetHeader("Authorization")

	if token == "" {
		err = errors.New("Authorization token is required")
		Forbidden(c, err)
		return
	}

	username, password, err := decodeToken(token)

	if err != nil {
		err = errors.New("Authorization token is required")
		Forbidden(c, err)
		return
	}

	if IsAuthorizedUser(token) == false {
		authenticatedUser, _ := goauth.AuthenticateUser(username, password)
		authorizedUser, _ := goauth.AuthorizeUser(authenticatedUser.Token)

		if authorizedUser.Authorize == true {
			AuthorizeUser(token)
		} else {
			err = errors.New("Authorization token is incorrect")
			Forbidden(c, err)
			return
		}
	}

	c.Next()
}
