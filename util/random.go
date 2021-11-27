package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/goombaio/namegenerator"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i:= 0; i< n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	return name 
}

// RandomBalance generates a random owner's balance
func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{USD, PHP}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email for unit tests
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}