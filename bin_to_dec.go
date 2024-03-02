package main

import (
	"fmt"
	"math"
)

func binaryToDecimal (n int) int {
	dec := 0
	i := 0
	for n != 0 {
		dec += (n % 10) * int(math.Pow(2, float64(i)))
		n /= 10
		i++
	}
	return dec
}

func main() {
	n := 111
	fmt.Println(binaryToDecimal(n))
	n = 1010
	fmt.Println(binaryToDecimal(n))
}
