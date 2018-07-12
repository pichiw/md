package md

import (
	"crypto/rand"
	"encoding/base64"
)

// RandID generates a random ID
func RandID() string {
	iddata := make([]byte, 10)
	_, err := rand.Read(iddata)
	if err != nil {
		panic(err) // this really shouldn't happen
	}
	return base64.RawURLEncoding.EncodeToString(iddata)
}
