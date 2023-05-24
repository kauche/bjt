package bjt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

var ErrInvalidToken = errors.New("invalid token")

// Token is a token that is encoded to and decoded from a base64 JSON string.
type Token[T any] struct {
	Source *T
}

// NewToken creates a new Token with the given source.
func NewToken[T any](source *T) *Token[T] {
	return &Token[T]{Source: source}
}

// Encode encodes the Token to a base64 string by following steps:
//  1. Marshal the Token.Source to a JSON object.
//  2. Encode it as base64.
func (p *Token[T]) Encode() (string, error) {
	jsonBytes, err := json.Marshal(p.Source)
	if err != nil {
		return "", fmt.Errorf("failed to marshal Token to a json object: %w", err)
	}

	return base64.StdEncoding.EncodeToString(jsonBytes), nil
}

// Decode decodes the given token string to a Token by following steps:
//  1. Decode the given token string as base64.
//  2. Unmarshal it as a JSON object and populate it to a Toekn with Token.Source that has been given as a generic type.
func Decode[T any](str string) (*Token[T], error) {
	base64Bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the token `%s`: %w, due to %w", str, ErrInvalidToken, err)
	}

	source := new(T)

	if err := json.Unmarshal(base64Bytes, source); err != nil {
		return nil, fmt.Errorf("failed to unmarshal the token `%s`: %w, due to %w", str, ErrInvalidToken, err)
	}

	return &Token[T]{Source: source}, nil
}
