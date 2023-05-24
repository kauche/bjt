package bjt_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/kauche/bjt"
)

type testTokenSource struct {
	ID    string `json:"id"`
	Order int    `json:"order"`
}

func TestToken(t *testing.T) {
	t.Parallel()

	s := &testTokenSource{
		ID:    "4bf5c56f-ba53-4b99-8a11-272e18ecb362",
		Order: 123,
	}

	et := bjt.NewToken(s)

	str, err := et.Encode()
	if err != nil {
		t.Errorf("failed to encode the token: %s\n", err)
		return
	}

	dt, err := bjt.Decode[testTokenSource](str)
	if err != nil {
		t.Errorf("failed to decode the token: %s\n", err)
		return
	}

	if dt.Source.ID != s.ID {
		t.Errorf("decoded token source ID is not matched with the original one: -%s +%s\n", dt.Source.ID, s.ID)
		return
	}

	if dt.Source.Order != s.Order {
		t.Errorf("decoded token source Order is not matched with the original one: -%d +%d\n", dt.Source.Order, s.Order)
		return
	}
}

func TestDecode_Error(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		token      string
		wantPrefix string
	}{
		"empty token": {
			token:      "",
			wantPrefix: "failed to unmarshal the token",
		},
		"invalid base64 token": {
			token:      "invalidtokenstring",
			wantPrefix: "failed to decode the token",
		},
	} {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			_, err := bjt.Decode[testTokenSource](test.token)
			if !errors.Is(err, bjt.ErrInvalidToken) {
				t.Errorf("error is not ErrInvalidToken: %s\n", err)
				return
			}

			if !strings.HasPrefix(err.Error(), test.wantPrefix) {
				t.Errorf("error message is not matched: -%s +%s\n", err.Error(), test.wantPrefix)
				return
			}
		})
	}
}

func TestEncode_Error(t *testing.T) {
	t.Parallel()

	ch := make(chan struct{}) // chan is not unmashalable to JSON
	token := bjt.NewToken[chan struct{}](&ch)
	_, err := token.Encode()
	if err == nil {
		t.Errorf("error is nil\n")
		return
	}

	if !strings.HasPrefix(err.Error(), "failed to marshal Token to a json object") {
		t.Errorf("error message is not matched: -%s +%s\n", err.Error(), "failed to marshal Token to a json object")
		return
	}
}
