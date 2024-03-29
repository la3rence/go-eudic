package go_eudic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	BaseURL   = "https://api.frdic.com"
	UserAgent = "/eusoft_eudic_en_mac/4.0.0/A4:83:E7:90:13:D5/explain/"
)

type EudicClient struct {
	BaseURL   string
	UserAgent string
	UserId    string
	Token     string
	client    *http.Client
	Response  *Response

	LoginService       *LoginService
	UserInfoService    *UserInfoService
	CheckInService     *CheckInService
	LastBookService    *LastBookService
	SyncReciteService  *SyncReciteService
	StartReciteService *StartReciteService
	AnswerCardService  *AnswerCardService
}

func NewEudicClientByPassword(username, password string) (*EudicClient, error) {
	client := &EudicClient{
		BaseURL:   BaseURL,
		UserAgent: UserAgent,
		client:    &http.Client{},
	}
	loginService := &LoginService{client: client}
	loginResponse, err := loginService.Login(username, password)
	if err != nil {
		return nil, err
	}
	return NewEudicClient(loginResponse.Userid, loginResponse.Token), nil
}

func NewEudicClient(userId, token string) *EudicClient {
	client := &EudicClient{
		BaseURL:   BaseURL,
		UserAgent: UserAgent,
		UserId:    userId,
		Token:     token,
		client:    &http.Client{},
	}

	client.LoginService = &LoginService{client: client}
	client.UserInfoService = &UserInfoService{client: client}
	client.CheckInService = &CheckInService{client: client}
	client.LastBookService = &LastBookService{client: client}
	client.SyncReciteService = &SyncReciteService{client: client}
	client.StartReciteService = &StartReciteService{client: client}
	client.AnswerCardService = &AnswerCardService{client: client}
	return client
}

func (eudic *EudicClient) NewRequest(method string, uri string, body string) (*http.Request, error) {
	req, err := http.NewRequest(method, eudic.BaseURL+uri, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("User-Agent", eudic.UserAgent)
	req.Header.Set("Origin", "http://dict.eudic.net")
	req.Header.Set("Referer", "http://dict.eudic.net")
	return req, nil
}

// v: response
func (eudic *EudicClient) Do(req *http.Request, v interface{}) (*Response, error) {
	httpResponse, err := eudic.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := httpResponse.Body.Close()
		if err != nil {
			fmt.Printf("error closing body: %+v", err)
		}
	}()

	err = CheckHTTPResponse(httpResponse)
	if err != nil {
		return nil, err
	}
	res := new(Response)
	bodyBytes, _ := ioutil.ReadAll(httpResponse.Body)
	tempStr := string(bodyBytes)
	res.BodyStrPtr = &tempStr
	if v != nil {
		res.Data = v
		err = json.Unmarshal(bodyBytes, res.Data)
		eudic.Response = res
	}
	return res, nil
}

func CheckHTTPResponse(httpResponse *http.Response) error {
	if httpResponse.StatusCode == http.StatusOK {
		return nil
	}
	if httpResponse.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("认证失败 [%d]", httpResponse.StatusCode)
	}
	return fmt.Errorf("无效的请求 [%d]", httpResponse.StatusCode)

}

type Response struct {
	HTTPResponse *http.Response
	BodyStrPtr   *string
	Data         interface{}
}

type RequestTokenBody struct {
	Token         string `json:"token"`
	UserId        string `json:"userid"`
	Cookie        string `json:"cookie"`
	EudicTimeZone int    `json:"eudic_timezone"`
	Lang          string `json:"lang"`
}
