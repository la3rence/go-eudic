package go_eudic

import "net/http"

var (
	getCheckInInfoURI = "/route/recite/getcheckininfo"
)

type CheckInInfoService struct {
	client *EudicClient
}

func (service *CheckInInfoService) GetCheckInInfo() (*CheckInInfoResponse, *string, error) {
	checkInInfoRequestBody := CheckInInfoRequest{
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

type CheckInInfoRequest struct {
	Cookie        string `json:"cookie"`
	EudicTimezone int64  `json:"eudic_timezone"`
	Lang          string `json:"lang"`
	Token         string `json:"token"`
	UseReciteUa   bool   `json:"use_recite_ua"`
	Userid        string `json:"userid"`
}

type CheckInInfoResponse struct {
	CheckinDate string `json:"checkin_date"`
	Continuous  int64  `json:"continuous"`
	Count       int64  `json:"count"`
	Goldcount   int64  `json:"goldcount"`
	Ischeckin   bool   `json:"ischeckin"`
}
