package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomString(n int) string {
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
// random owner name
func RandomOwner() string {
	return RandomString(6)
}

// random amount
func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

// random currency
func RandomCurrency() string {
	currencies := []string{"USD", "KHR", "THB"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
