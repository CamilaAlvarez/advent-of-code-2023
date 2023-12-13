package math

import "math"

func IntPow(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
