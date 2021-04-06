package go_eudic

import "os"

var (
	client *EudicClient
)

func setup() {
	client = NewEudicClient(
		os.Getenv("EUDIC_USERID"),
		os.Getenv("EUDIC_TOKEN"),
	)
}
