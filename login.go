package go_eudic

import (
	"net/http"
)

var authorizeURI = "/api/v2/auth/authorize"

// I don't know how it generated. If 401 returned due to this variable, please submit an issue
var authorizeString = "QYN eyJjYW1wYWlnbiI6ZmFsc2UsImZsIjowLjAsImxjIjowLCJ0IjoiQUJJTVRZek56VTROelF3T1E9PSIsInRva" +
	"2VuIjoiIiwidXJsc2lnbiI6Ik44ZjFtYUFHbmR3VUpscVljVWt5ajlhSjNqOD0iLCJ1c2VyaWQiOiIiLCJ2X2RpY3QiOjF9Cg=="

type LoginService struct {
	client *EudicClient
}

func (service *LoginService) Login(username, password string) (*LoginResponse, error) {
	loginRequestBody := &LoginRequest{
		Password: password,
		Username: username,
	}

	loginRequestString, err := StructToString(loginRequestBody)
	if err != nil {
		return nil, err
	}

	loginRequest, err := service.client.NewRequest(http.MethodPost, authorizeURI, loginRequestString)
	if err != nil {
		return nil, err
	}
	loginRequest.Header.Set("Authorization", authorizeString)
	loginResponse := new(LoginResponse)
	_, err = service.client.Do(loginRequest, loginResponse)
	if err != nil {
		return nil, err
	}
	return loginResponse, nil
}

type LoginRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginResponse struct {
	CreationDate            string      `json:"creation_date"`
	Expirein                int64       `json:"expirein"`
	HasOldPassword          bool        `json:"has_old_password"`
	LastPasswordChangedDate interface{} `json:"last_password_changed_date"`
	OpenidDesc              interface{} `json:"openid_desc"`
	OpenidType              interface{} `json:"openid_type"`
	Profile                 struct {
		Email        string      `json:"email"`
		Gender       string      `json:"gender"`
		Nickname     string      `json:"nickname"`
		Password     interface{} `json:"password"`
		Vocabularies struct{}    `json:"vocabularies"`
	} `json:"profile"`
	RedirectURL interface{} `json:"redirect_url"`
	Roles       interface{} `json:"roles"`
	Token       string      `json:"token"`
	Userid      string      `json:"userid"`
	Username    string      `json:"username"`
}
