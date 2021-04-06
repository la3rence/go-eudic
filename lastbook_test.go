package go_eudic

import (
	"fmt"
	"testing"
)

func TestLastBookService_GetLastBook(t *testing.T) {
	setup()
	lastBookRes, lastBookString, err := client.LastBookService.GetLastBook()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*lastBookString)
	if lastBookRes.Meta.Userid != client.UserId {
		t.Errorf("Reponse from server doesn't match user info, expected userid is %s, bug %s got",
			client.UserId, lastBookRes.Meta.Userid)
	}
}
