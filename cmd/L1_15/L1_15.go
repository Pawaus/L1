package L1_15

import (
	"fmt"
	"math/rand"
)

var justString string

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func createHugeString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func someFunc() string {
	return createHugeString(1 << 10)[:100]
}

func Fifteen() {
	fmt.Println(someFunc())
}
