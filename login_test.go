package go_eudic

import (
	"os"
	"testing"
)

func TestLoginService_Login(t *testing.T) {
	setup(t)
	_, err := client.LoginService.Login(
		os.Getenv("EUDIC_USERNAME"),
		os.Getenv("EUDIC_PASSWORD"))
	if err != nil {
		t.Error(err)
	}
}
