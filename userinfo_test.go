package go_eudic

import (
	"fmt"
	"testing"
)

func TestUserInfoService_GetUserInfo(t *testing.T) {
	setup(t)
	responseUserInfo, _, err := client.UserInfoService.GetUserInfo()
	if err != nil {
		t.Error(err)
	}
	if responseUserInfo.Userid != client.UserId {
		t.Errorf("Reponse from server doesn't match user info, expected userid is %s, bug %s got",
			client.UserId, responseUserInfo.Userid)
	}
	fmt.Printf("%+v\n", responseUserInfo)
}
