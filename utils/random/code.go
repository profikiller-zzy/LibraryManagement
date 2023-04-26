package random

import (
	"math/rand"
	"time"
)

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)
	setLen := len(charSet)
	for i := range code {
		code[i] = charSet[rand.Intn(100000000)%setLen]
	}
	return string(code)
}
