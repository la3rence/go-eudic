package go_eudic

import (
	"testing"
	"time"
)

func TestCheckInInfoService_GetCheckInInfo(t *testing.T) {
	setup(t)
	checkInBody, _, err := client.CheckInService.GetCheckInInfo()
	if err != nil {
		t.Error(err)
	}
	today := time.Now().Format(`2006-01-02`)
	if checkInBody.CheckinDate != today {
		t.Errorf("Wrong checkin info's date: %s expected but %s got", today, checkInBody.CheckinDate)
	}
}

func TestCheckInService_CheckIn(t *testing.T) {
	setup(t)
	checkInResponse, err := client.CheckInService.CheckIn()
	if err != nil {
		t.Error(err)
	}
	if !checkInResponse.Ischeckin {
		t.Errorf("Fail to checkin: the checkin state is `false`, expected `true`")
	}
}
