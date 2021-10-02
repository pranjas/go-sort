package randUtil

import (
	"math/rand" /*For random numbers and dynamic test arrays*/
	"time"      /*For seeding random generator*/
)

var randomSource rand.Source
var randomGenerator *rand.Rand

func init() {
	randomSource = rand.NewSource(time.Now().UnixNano())
	randomGenerator = rand.New(randomSource)
}

func GenerateRandomInt(maxValue int) int {
	return randomGenerator.Intn(maxValue)
}

func GenerateRandomIntIncludingNegatives(maxValue int) int {
	return -maxValue + randomGenerator.Intn(2*maxValue+1)
}
