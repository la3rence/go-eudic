package go_eudic

import (
	"net/http"
)

var startReciteURI = "/route/recite/startrecite"

type StartReciteService struct {
	client *EudicClient
}

// 开始背诵词汇，bookId 为纯数字字符串, 并非 LastBook 返回的 Meta.BookId, 调用时正则取其中的数字，参考测试
func (service *StartReciteService) StartRecite(bookId, bookName string) (*ReciteResponse, error) {
	requestBody := StartReciteRequest{
		BookID:        bookId,
		BookName:      bookName,
		DailyCount:    30,
		EudicTimezone: 8,
		Lang:          "en",
		ListType:      2,
		Token:         service.client.Token,
		UseReciteUa:   false,
		Userid:        service.client.UserId,
	}
	startReciteRequestString, err := StructToString(requestBody)
	if err != nil {
		return nil, err
	}
	startReciteRequest, err := service.client.NewRequest(http.MethodPost, startReciteURI, startReciteRequestString)
	startReciteResponse := new(ReciteResponse)
	_, err = service.client.Do(startReciteRequest, startReciteResponse)
	if err != nil {
		return nil, err
	}
	//fmt.Println(*response.BodyStrPtr)
	return startReciteResponse, nil
}

type StartReciteRequest struct {
	BookID        string `json:"book_id"`
	BookName      string `json:"book_name"`
	BookUUID      string `json:"book_uuid"`
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

type ReciteResponse struct {
	Card struct {
		AddTime       int64  `json:"addTime"`
		Answer        string `json:"answer"`
		AnswerHistory []struct {
			DueTimeInterval int64   `json:"dueTimeInterval"`
			EaseFactor      float64 `json:"easeFactor"`
			EaseLevel       int64   `json:"easeLevel"`
			Level           int64   `json:"level"`
			Meta            string  `json:"meta"`
			ReciteDuration  int64   `json:"reciteDuration"`
			ReciteTime      int64   `json:"reciteTime"`
		} `json:"answerHistory"`
		CandidateExplains []struct {
			Answer   string `json:"answer"`
			Question string `json:"question"`
		} `json:"candidateExplains"`
		CardID                 int64       `json:"card_id"`
		CurrentReciteEndTime   int64       `json:"currentReciteEndTime"`
		CurrentReciteStartTime int64       `json:"currentReciteStartTime"`
		DueTime                int64       `json:"dueTime"`
		EaseFactor             float64     `json:"easeFactor"`
		FirstReciteTime        int64       `json:"firstReciteTime"`
		LastDueTime            int64       `json:"lastDueTime"`
		LastEaseFactor         float64     `json:"lastEaseFactor"`
		Level                  int64       `json:"level"`
		LocalTimestamp         int64       `json:"local_timestamp"`
		LocalUpdate            bool        `json:"local_update"`
		Question               string      `json:"question"`
		Rating                 int64       `json:"rating"`
		ResInfo                interface{} `json:"resInfo"`
		Status                 int64       `json:"status"`
		Timezone               int64       `json:"timezone"`
		Tombstone              bool        `json:"tombstone"`
		TotalBrowseTime        int64       `json:"totalBrowseTime"`
		TotalReciteTime        int64       `json:"totalReciteTime"`
		Unit                   int64       `json:"unit"`
	} `json:"card"`
	Explain struct {
		Phonukurl string `json:"phonukurl"`
		Phonusurl string `json:"phonusurl"`
		Word      string `json:"word"`
	} `json:"explain"`
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
	TaskFinished  bool   `json:"taskFinished"`
	TaskInfo      string `json:"taskInfo"`
	TodayProgress struct {
		PendingDueCardCount int64 `json:"pendingDueCardCount"`
		PendingNewCardCount int64 `json:"pendingNewCardCount"`
		TodayFinishedCount  int64 `json:"todayFinishedCount"`
		TodayTouchedCount   int64 `json:"todayTouchedCount"`
	} `json:"todayProgress"`
}
