package utils

import (
	"math/rand"
	"strings"
)

func GenerateStringDomainName(N int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	r := make([]string, N)
	for i := 0; i < N; i++ {
		c := chars[rand.Intn(len(chars))]
		r[i] = string(c)
	}

	return strings.Join(r, "")
}
