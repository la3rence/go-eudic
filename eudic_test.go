package go_eudic

import (
	"fmt"
	"os"
	"testing"
)

var client *EudicClient

func setup(t *testing.T) {
	tempClient, err := NewEudicClientByPassword(
		os.Getenv("EUDIC_USERNAME"),
		os.Getenv("EUDIC_PASSWORD"),
	)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tempClient.Token)
	client = tempClient
}
