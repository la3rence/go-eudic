package go_eudic

import (
	"net/http"
)

var (
	getLastBookURI = "/route/recite/getlastbook"
)

type LastBookService struct {
	client *EudicClient
}

func (service *LastBookService) GetLastBook() (*LastBodyResponse, *string, error) {
	request := LastBookRequest{
		Cookie:      "",
		DailyCount:  30,
		Lang:        "en",
		Token:       service.client.Token,
		UseReciteUa: false,
		Userid:      service.client.UserId,
	}
	lastBookRequestJSON, err := StructToString(request)
	if err != nil {
		return nil, nil, err
	}
	getLastBookReq, err := service.client.NewRequest(http.MethodPost, getLastBookURI, lastBookRequestJSON)
	lastBookResponse := new(LastBodyResponse)
	response, err := service.client.Do(getLastBookReq, lastBookResponse)
	if err != nil {
		return nil, nil, err
	}
	return lastBookResponse, response.BodyStrPtr, nil
}

type LastBookRequest struct {
	Cookie      string `json:"cookie"`
	DailyCount  int64  `json:"daily_count"`
	Lang        string `json:"lang"`
	Token       string `json:"token"`
	UseReciteUa bool   `json:"use_recite_ua"`
	Userid      string `json:"userid"`
}

type LastBodyResponse struct {
	BookName string `json:"bookName"`
	Meta     struct {
		Bookid    string `json:"bookid"`
		Lang      string `json:"lang"`
		Meta      string `json:"meta"`
		Timestamp string `json:"timestamp"`
		Userid    string `json:"userid"`
		Version   string `json:"version"`
	} `json:"meta"`
	StatisticData struct {
		AverageDailyLearningCardCount float64 `json:"averageDailyLearningCardCount"`
		DailyHistory                  []struct {
			Day             int64   `json:"day"`
			LevelDifference int64   `json:"levelDifference"`
			NewCount        int64   `json:"newCount"`
			Progress        float64 `json:"progress"`
			ReciteCount     int64   `json:"reciteCount"`
			ReciteDuration  int64   `json:"reciteDuration"`
		} `json:"dailyHistory"`
		LearningCardCount int64   `json:"learningCardCount"`
		MatureCardCount   int64   `json:"matureCardCount"`
		NewCardCount      int64   `json:"newCardCount"`
		Progress          float64 `json:"progress"`
		TotalCardCount    int64   `json:"totalCardCount"`
	} `json:"statisticData"`
	TodayProgress struct {
		PendingDueCardCount int64 `json:"pendingDueCardCount"`
		PendingNewCardCount int64 `json:"pendingNewCardCount"`
		TodayFinishedCount  int64 `json:"todayFinishedCount"`
		TodayTouchedCount   int64 `json:"todayTouchedCount"`
	} `json:"todayProgress"`
}
