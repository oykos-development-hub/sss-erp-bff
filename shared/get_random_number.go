package shared

import (
	"math/rand"
	"time"
)

var (
	// A local instance of *rand.Rand, initialized once
	localRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GetRandomNumber() int {
	// Generate a random integer between 100 and 100,000
	randomInt := localRand.Intn(100000-100+1) + 100

	return randomInt
}
