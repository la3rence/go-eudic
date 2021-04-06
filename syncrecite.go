package go_eudic

import (
	"net/http"
)

var (
	syncReciteURI = "/route/recite/syncrecite"
)

type SyncReciteService struct {
	client *EudicClient
}

// Sync the learning records
func (service *SyncReciteService) SyncRecite(bookId, bookName string) (bool, error) {
	syncReciteRequestBody := SyncReciteRequest{
		BookID:      bookId,
		BookName:    bookName,
		BookUUID:    "",
		Cookie:      "",
		Lang:        "en",
		ListType:    2, // dont know what this mean
		Token:       service.client.Token,
		UseReciteUa: false,
		Userid:      service.client.UserId,
	}
	syncReciteRequestString, err := StructToString(syncReciteRequestBody)
	if err != nil {
		return false, err
	}
	syncReciteRequest, err := service.client.NewRequest(http.MethodPost, syncReciteURI, syncReciteRequestString)
	success := new(string)
	response, err := service.client.Do(syncReciteRequest, success)
	if err != nil {
		return false, err
	}
	if *response.BodyStrPtr == "success" {
		return true, nil
	} else {
		return false, nil
	}
}

type SyncReciteRequest struct {
	BookID      string `json:"book_id"`
	BookName    string `json:"book_name"`
	BookUUID    string `json:"book_uuid"`
	Cookie      string `json:"cookie"`
	Lang        string `json:"lang"`
	ListType    int64  `json:"list_type"`
	Token       string `json:"token"`
	UseReciteUa bool   `json:"use_recite_ua"`
	Userid      string `json:"userid"`
}
