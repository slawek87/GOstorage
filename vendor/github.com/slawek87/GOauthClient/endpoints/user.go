package endpoints


type User struct {
	Username 	string
	Password 	string
}

type Token struct {
	Token       string
	Authorize 	bool
}

type Password struct {
	Result  	bool
}

func (user *User) RegisterUser(username string, password string) (string, map[string]string) {
	const REGISTER_USER = "/api/v1/user/register"

	formData := map[string]string {
		"Username": username,
		"Password": password,
	}

	return REGISTER_USER, formData
}

func (user *User) AuthenticateUser(username string, password string) (string, map[string]string) {
	const AUTHENTICATE_USER = "/api/v1/user/authentication"

	formData := map[string]string {
		"Username": username,
		"Password": password,
	}

	return AUTHENTICATE_USER, formData
}

func (user *User) AuthorizeUser(token string) (string, map[string]string) {
	const AUTHORIZATION_USER = "/api/v1/user/authorization"

	formData := map[string]string {
		"Token": token,
	}

	return AUTHORIZATION_USER, formData
}

func (user *User) ResetUserPassword(username string, password string) (string, map[string]string) {
	const PASSWORD_RESET = "/api/v1/user/password/reset"

	formData := map[string]string {
		"Username": username,
		"Password": password,
	}

	return PASSWORD_RESET, formData
}