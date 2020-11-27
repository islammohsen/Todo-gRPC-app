package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generate a random integer Between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

//RandomString generates a random string of length n
func RandomString(n int) string {

	runes := make([]rune, n)

	for i := 0; i < n; i++ {
		runes[i] = 'a' + rune(RandomInt(0, 26))
	}

	return string(runes)
}
