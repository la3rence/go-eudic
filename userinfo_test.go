package go_eudic

import (
	"testing"
)

func TestUserInfoService_GetUserInfo(t *testing.T) {
	setup()
	responseUserInfo, _, err := client.UserInfoService.GetUserInfo()
	if err != nil {
		t.Error(err)
	}
	if responseUserInfo.Userid != client.UserId {
		t.Errorf("Reponse from server doesn't match user info, expected userid is %s, bug %s got",
			client.UserId, responseUserInfo.Userid)
	}
}
