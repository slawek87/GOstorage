package authorization

import (
	"github.com/slawek87/GOauthClient/client"
	"fmt"
)

var Settings = map[string]string {
	"HOST": "0.0.0.0",
	"PROTOCOL": "http",
	"PORT": "8090",
	"SERVICE_TOKEN": "580f4525-a97a-4ce8-8b5a-d4cf359771ae",
	"SERVICE_LOGIN": "GOblog",
}

func Example() {
	settings := client.Client{Settings: Settings}
	goauth := client.GOauth{Client: settings}

	USER := "slawek1987"
	PASSWORD := "k1k2k3k4"

	goauth.RegisterUser(USER, PASSWORD)
	user, _ := goauth.AuthenticateUser(USER, PASSWORD)
	token, _ := goauth.AuthorizeUser(user.Token)

	client.AuthorizeUser(token.Token)

	fmt.Println(client.IsAuthorizedUser(token.Token))
}
