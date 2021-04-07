package go_eudic

import (
	"fmt"
	"testing"
)

func TestAnswerCardService_AnswerCard(t *testing.T) {
	setup(t)
	book, _, _ := client.LastBookService.GetLastBook()
	bookId := GetTheFirstNumberFromString(book.Meta.Bookid)

	startRecite, err := client.StartReciteService.StartRecite(bookId, book.BookName)
	if err != nil {
		t.Error(err)
	}
	cardId := startRecite.Card.CardID
	question := startRecite.Card.Question
	fmt.Printf("ID: %d\n单词: %s\n", cardId, question)
	// 5: 熟悉 2: 模糊 0: 完全不认识
	_, err = client.AnswerCardService.AnswerCard(bookId, book.BookName, cardId, 2)
	if err != nil {
		t.Error(err)
	}
}
