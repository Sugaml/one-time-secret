package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())

}

//RandomInt generates a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomContent generate a random owner name
func RandomContent() string {
	return RandomString(100)
}
func RandomUser() string {
	return RandomString(6)
}
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
func RandomName() string {
	return RandomString(7)
}
