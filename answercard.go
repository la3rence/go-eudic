package go_eudic

import "net/http"

var answerCardURI = "/route/recite/answercard"

type AnswerCardService struct {
	client *EudicClient
}

func (service *AnswerCardService) AnswerCard(bookId, bookName string, cardId, answerEase int64) (*ReciteResponse, error) {
	// answerEase: 5 认识 2 模糊 不认识 0
	requestBody := AnswerCardRequest{
		AnswerEase:    answerEase,
		BookID:        bookId,
		BookName:      bookName,
		CardID:        cardId,
		DailyCount:    30,
		EudicTimezone: 8,
		Lang:          "en",
		ListType:      2,
		Token:         service.client.Token,
		Userid:        service.client.UserId,
	}
	answerRequestString, err := StructToString(requestBody)
	if err != nil {
		return nil, err
	}
	answerRequest, err := service.client.NewRequest(http.MethodPost, answerCardURI, answerRequestString)
	answerResponse := new(ReciteResponse)
	_, err = service.client.Do(answerRequest, answerResponse)
	if err != nil {
		return nil, err
	}
	return answerResponse, nil
}

type AnswerCardRequest struct {
	AnswerEase    int64  `json:"answer_ease"`
	BookID        string `json:"book_id"`
	BookName      string `json:"book_name"`
	BookUUID      string `json:"book_uuid"`
	CardID        int64  `json:"card_id"`
	Cookie        string `json:"cookie"`
	DailyCount    int64  `json:"daily_count"`
	EudicTimezone int64  `json:"eudic_timezone"`
	Lang          string `json:"lang"`
	ListType      int64  `json:"list_type"`
	Meta          string `json:"meta"`
	Token         string `json:"token"`
	UseReciteUa   bool   `json:"use_recite_ua"`
	Userid        string `json:"userid"`
}
