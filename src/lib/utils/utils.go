package utils

import (
	"math/rand"
)

func GetRandom(min int, max int) int {
	return min + rand.Intn(max-min+1)
}
