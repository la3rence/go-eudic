package go_eudic

import (
	"fmt"
	"testing"
	"time"
)

func TestCheckInInfoService_GetCheckInInfo(t *testing.T) {
	setup(t)
	checkInBody, checkInString, err := client.CheckInInfoService.GetCheckInInfo()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*checkInString)
	today := time.Now().Format(`2006-01-02`)
	if checkInBody.CheckinDate != today {
		t.Errorf("Wrong checkin date: %s expected but %s got", today, checkInBody.CheckinDate)
	}
}
