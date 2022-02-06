package generator

import (
	"crypto/rand"
	"io"
)

func EncodeToString(max int) (string, error) {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), err
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
