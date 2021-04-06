package go_eudic

import (
	"net/http"
)

var (
	getUserInfoURI = "/route/recite/getuserinfo"
)

type UserInfoService struct {
	client *EudicClient
}

func (service *UserInfoService) GetUserInfo() (*UserInfo, *string, error) {
	requestBody := RequestTokenBody{
		Token:         service.client.Token,
		UserId:        service.client.UserId,
		Cookie:        "",
		EudicTimeZone: 8,
		Lang:          "en",
	}
	bodyString, err := StructToString(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request, err := service.client.NewRequest(http.MethodPost, getUserInfoURI, bodyString)
	userInfo := new(UserInfo)
	infoResponse, err := service.client.Do(request, userInfo)
	if err != nil {
		return nil, nil, err
	}
	return userInfo, infoResponse.BodyStrPtr, nil
}

type UserInfo struct {
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
