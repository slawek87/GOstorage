package auth

import "github.com/slawek87/GOauthClient/client"


var Settings = map[string]string {
	"HOST": "0.0.0.0",
	"PROTOCOL": "http",
	"PORT": "8090",
	"SERVICE_TOKEN": "fb008237-c46a-4823-b027-c2ea2f49340d",
	"SERVICE_LOGIN": "GOstorage",
}

func GOauth() client.GOauth {
	settings := client.Client{Settings: Settings}
	goauth := client.GOauth{Client: settings}

	return goauth
}