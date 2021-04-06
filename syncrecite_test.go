package go_eudic

import (
	"fmt"
	"testing"
)

func TestSyncReciteService_SyncRecite(t *testing.T) {
	setup()
	book, _, _ := client.LastBookService.GetLastBook()
	bookId := GetTheFirstNumberFromString(book.Meta.Bookid)
	syncRecite, err := client.SyncReciteService.SyncRecite(bookId, book.BookName)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(syncRecite)
	if !syncRecite {
		t.Errorf("Sync recite fail: %+v\n", syncRecite)
	}
}
