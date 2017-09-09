package client

import (
	"gopkg.in/resty.v0"
	"github.com/slawek87/GOauthClient/endpoints"
	"reflect"
)

type Client struct {
	Settings    map[string]string
}

type GOauth struct {
	Client		Client
}

func (goAuth *GOauth) RegisterService(name string) (endpoints.ServiceEndpoint, map[string]interface{}) {
	serviceEndpoint := endpoints.ServiceEndpoint{}

	url, formData := serviceEndpoint.RegisterService(name)
	data, _ := goAuth.Client.postRequest(url, formData, &serviceEndpoint)

	return serviceEndpoint, data
}

func (goAuth *GOauth) RegisterUser(username string, password string) (endpoints.User, map[string]interface{}) {
	userEndpoint := endpoints.User{}

	url, formData := userEndpoint.RegisterUser(username, password)
	data, _ := goAuth.Client.postRequest(url, formData, &userEndpoint)

	return userEndpoint, data
}

func (goAuth *GOauth) ResetUserPassword(username string, password string) (endpoints.User, map[string]interface{}) {
	passwordEndpoint := endpoints.Password{}
	userEndpoint := endpoints.User{}

	url, formData := userEndpoint.ResetUserPassword(username, password)
	data, _ := goAuth.Client.postRequest(url, formData, &passwordEndpoint)

	return userEndpoint, data
}

func (goAuth *GOauth) AuthenticateUser(username string, password string) (endpoints.Token, map[string]interface{}) {
	tokenEndpoint := endpoints.Token{}
	userEndpoint := endpoints.User{}

	url, formData := userEndpoint.AuthenticateUser(username, password)
	data, _ := goAuth.Client.postRequest(url, formData, &tokenEndpoint)

	if tokenEndpoint.Token != "" {
		tokenEndpoint.Authorize = true
	}

	return tokenEndpoint, data
}

func (goAuth *GOauth) AuthorizeUser(token string) (endpoints.Token, map[string]interface{}) {
	tokenEndpoint := endpoints.Token{}
	userEndpoint := endpoints.User{}

	url, formData := userEndpoint.AuthorizeUser(token)
	data, _ := goAuth.Client.postRequest(url, formData, &tokenEndpoint)

	if tokenEndpoint.Token != "" {
		tokenEndpoint.Authorize = true
	}

	tokenEndpoint.Token = token

	return tokenEndpoint, data
}

func (client *Client) resty() *resty.Request {
	settings := client.Settings
    request := resty.R()
	request.SetHeader("Content-Type", "application/json")
	request.SetBasicAuth(settings["SERVICE_LOGIN"], settings["SERVICE_TOKEN"])

	return request
}

func (client *Client) mapStructure(data map[string]interface{}, result interface{}) interface{} {
	elements := reflect.ValueOf(result).Elem()

	for key, value := range data {
		getValue := elements.FieldByName(key)

		if getValue.IsValid() {
			getValue.Set(reflect.ValueOf(value))
		}
	}

	return result
}

func (client *Client) postRequest(url string, formData map[string]string, result interface{}) (map[string]interface{}, interface{}) {
	var data map[string]interface{}

	client.resty().SetFormData(formData).SetResult(&data).Post(client.GetURL(url))
	client.mapStructure(data, result)

	return data, result
}
