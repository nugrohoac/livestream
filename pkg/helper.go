package pkg

import (
	"encoding/base64"
	"time"
)

// EncodeCursor takes a interface and encode its string representation
// using Base64 encoding by cursor type
func EncodeCursor(cursor interface{}) (string, error) {
	byt, err := cursor.(time.Time).MarshalText()
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(byt), nil
}

// DecodeCursor takes a string and decode it using Base64 encoding to convert it
// into time
func DecodeCursor(cursor string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return time.Time{}, err
	}

	return time.Parse(time.RFC3339, string(byt))
}
