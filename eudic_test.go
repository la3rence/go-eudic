package go_eudic

import (
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
	client = tempClient
}
