package settings

import "github.com/slawek87/GOauthClient/client"

var Settings = map[string]string {
	"HOST": "0.0.0.0",
	"PROTOCOL": "http",
	"PORT": "8090",
	"SERVICE_TOKEN": "580f4525-a97a-4ce8-8b5a-d4cf359771ae",
	"SERVICE_LOGIN": "GOblog",
}

func GOauth() client.GOauth {
	settings := client.Client{Settings: Settings}
	goauth := client.GOauth{Client: settings}

	return goauth
}