package go_eudic

import (
	"fmt"
	"testing"
)

func TestLoginService_Login(t *testing.T) {
	setup(t)
	loginResponse, err := client.LoginService.Login("lonor@live.com", "EUDIC1412")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(loginResponse.Token)
	fmt.Println(loginResponse.Userid)
}
