# bjt

`bjt` (stannds for Base64 JSON Token) is a Go library that provides a simple way to encode and decode tokens using a base64 encoded JSON format.

## Supported Go version

1.20 or higher.

## Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/kauche/bjt"
)

type myTokenSource struct {
	ID     string `json:"id"`
	Number int    `json:"number"`
}

func main() {
	source := &myTokenSource{
		ID:     "124b473d-6079-46c4-b8bd-3824c93cef32",
		Number: 123,
	}

	token := bjt.NewToken(source)

	tokenStr, err := token.Encode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode: %s\n", err)
		os.Exit(1)
	}

	// This prints `eyJpZCI6IjEyNGI0NzNkLTYwNzktNDZjNC1iOGJkLTM4MjRjOTNjZWYzMiIsIm51bWJlciI6MTIzfQ==`
	// that is the base64 encoded myTokenSource JSON object: `{"id":"124b473d-6079-46c4-b8bd-3824c93cef32","number":123}`
	fmt.Println(tokenStr)

	decodedToken, err := bjt.Decode[myTokenSource](tokenStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode: %s\n", err)
		os.Exit(1)
	}

	// This prints `ID:124b473d-6079-46c4-b8bd-3824c93cef32 Number:123`
	fmt.Printf("ID:%s Number:%d\n", decodedToken.Source.ID, decodedToken.Source.Number)
}
```

See https://pkg.go.dev/github.com/kauche/bjt for more details.
