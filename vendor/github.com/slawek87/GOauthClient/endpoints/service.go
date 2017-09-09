package endpoints

type ServiceEndpoint struct {
    Name 	string
	Token 	string
}

// Method returns url for register service endpoint and body needed to send post request.
func (service *ServiceEndpoint) RegisterService(name string) (string, map[string]string) {
	const REGISTER_SERVICE = "/api/v1/service/register"

	formData := map[string]string {
		"Name": name,
	}

	return REGISTER_SERVICE, formData
}

