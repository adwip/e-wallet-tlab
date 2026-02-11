package utils

import (
	"math/rand/v2"
	"strconv"
)

func GenerateAccountNumber() string {
	var result = strconv.Itoa(rand.IntN(9) + 1)

	for i := 0; i < 7; i++ {
		result += strconv.Itoa(rand.IntN(10))
	}

	return result
}
