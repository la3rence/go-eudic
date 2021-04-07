# EUDIC SDK for Go

## REST endpoint

This SDK only supports for the part of reciting with the API host: `https://api.frdic.com`

General HTTP Headers:

- Content-Type: application/json;charset=utf-8
- User-Agent: /eusoft_eudic_en_mac/4.0.2/A4:83:E7:90:13:5D/explain/

```shell
export EUDIC_USERNAME="change me"
export EUDIC_PASSWORD="change me"
go test ./...
```

## Usage

```shell
go get -u github.com/Lonor/go-eudic
```

```go
package main

import (
	eudic "github.com/Lonor/go-eudic"
)

func main() {
	client, _ := eudic.NewEudicClientByPassword(
		"mail address",
		"your password",
	)
	book, _, _ := client.LastBookService.GetLastBook()
}
```

## LICENSE

This project is under the [GPLv3](https://www.gnu.org/licenses/gpl-3.0.html).