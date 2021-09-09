package utils

import (
	"math"
)

func IsPrime(number int) bool {
	if number <= 2 {
		return true
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
