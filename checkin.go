package go_eudic

import (
	"net/http"
)

var (
	checkInURI        = "/route/recite/checkin"
	getCheckInInfoURI = "/route/recite/getcheckininfo"
)

type CheckInService struct {
	client *EudicClient
}

func (service *CheckInService) CheckIn() (*CheckInResponse, error) {
	checkInRequestBody := CheckInRequest{
		Cookie:        "",
		EudicTimezone: 8,
		Lang:          "en",
		Token:         service.client.Token,
		UseReciteUa:   false,
		Userid:        service.client.UserId,
	}
	checkInRequestString, err := StructToString(checkInRequestBody)
	if err != nil {
		return nil, err
	}
	checkInRequest, err := service.client.NewRequest(http.MethodPost, checkInURI, checkInRequestString)
	checkInResponse := new(CheckInResponse)
	_, err = service.client.Do(checkInRequest, checkInResponse)
	if err != nil {
		return nil, err
	}
	return checkInResponse, nil
}

func (service *CheckInService) GetCheckInInfo() (*CheckInInfoResponse, *string, error) {
	checkInInfoRequestBody := CheckInRequest{
		Cookie:        "",
		EudicTimezone: 8,
		Lang:          "en",
		Token:         service.client.Token,
		UseReciteUa:   false,
		Userid:        service.client.UserId,
	}

	checkInInfoRequestString, err := StructToString(checkInInfoRequestBody)
	if err != nil {
		return nil, nil, err
	}
	checkInRequest, err := service.client.NewRequest(http.MethodPost, getCheckInInfoURI, checkInInfoRequestString)
	checkInInfoResponse := new(CheckInInfoResponse)
	response, err := service.client.Do(checkInRequest, checkInInfoResponse)
	if err != nil {
		return nil, nil, err
	}
	return checkInInfoResponse, response.BodyStrPtr, nil
}

type CheckInRequest struct {
	Cookie        string `json:"cookie"`
	EudicTimezone int64  `json:"eudic_timezone"`
	Lang          string `json:"lang"`
	Token         string `json:"token"`
	UseReciteUa   bool   `json:"use_recite_ua"`
	Userid        string `json:"userid"`
}

type CheckInResponse struct {
	CheckinDate string `json:"checkin_date"`
	Continuous  int64  `json:"continuous"`
	Count       int64  `json:"count"`
	Goldcount   int64  `json:"goldcount"`
	Ischeckin   bool   `json:"ischeckin"`
	Reward      int64  `json:"reward"`
}

type CheckInInfoResponse struct {
	CheckinDate string `json:"checkin_date"`
	Continuous  int64  `json:"continuous"`
	Count       int64  `json:"count"`
	Goldcount   int64  `json:"goldcount"`
	Ischeckin   bool   `json:"ischeckin"`
}
