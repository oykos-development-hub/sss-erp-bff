package shared

import (
	"math/rand"
	"time"
)

var counter = 0

func GetRandomNumber() int {
	// Initialize the random number generator with the current time as the seed
	rand.Seed(time.Now().UnixNano())
	// Generate a random integer between 100 and 100,000
	randomInt := rand.Intn(100000-100+1) + 100 + counter

	counter++

	return randomInt
}
