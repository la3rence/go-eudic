package go_eudic

import (
	"fmt"
	"testing"
)

func TestStartReciteService_StartRecite(t *testing.T) {
	setup()
	book, _, _ := client.LastBookService.GetLastBook()
	bookId := GetTheFirstNumberFromString(book.Meta.Bookid)
	startReciteRes, err := client.StartReciteService.StartRecite(bookId, book.BookName)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", startReciteRes)
}
