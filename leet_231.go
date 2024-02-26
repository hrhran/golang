package main

import (
	"fmt"
	"math"
)

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	logVal := math.Log2(float64(n))
	return math.Floor(logVal) == logVal
}

func main() {
	n := 64
	fmt.Println(isPowerOfTwo(n))
}
