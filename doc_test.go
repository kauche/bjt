package bjt_test

import (
	"fmt"
	"os"

	"github.com/kauche/bjt"
)

type myTokenSource struct {
	ID     string `json:"id"`
	Number int    `json:"number"`
}

func Example() {
	source := &myTokenSource{
		ID:     "124b473d-6079-46c4-b8bd-3824c93cef32",
		Number: 123,
	}

	token := bjt.NewToken(source)

	tokenStr, err := token.Encode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode: %s\n", err)
		return
	}

	// This prints `eyJpZCI6IjEyNGI0NzNkLTYwNzktNDZjNC1iOGJkLTM4MjRjOTNjZWYzMiIsIm51bWJlciI6MTIzfQ==`
	// that is the base64 encoded myTokenSource JSON object: `{"id":"124b473d-6079-46c4-b8bd-3824c93cef32","number":123}`
	fmt.Println(tokenStr)

	decodedToken, err := bjt.Decode[myTokenSource](tokenStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode: %s\n", err)
		return
	}

	// This prints `ID:124b473d-6079-46c4-b8bd-3824c93cef32 Number:123`
	fmt.Printf("ID:%s Number:%d\n", decodedToken.Source.ID, decodedToken.Source.Number)
}
