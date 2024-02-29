package main

import (
	"fmt"
	"math"
)

func isArmstrong(n int) bool {
	original := n
	count := 0
	i := 0
	for i = n; i != 0; count++ {
		i /= 10
	}
	sum := 0
	for i = n; i != 0; i /= 10 {
		sum += int(math.Pow(float64(i % 10), float64(count)))
	}
	return sum == original
}

func main() {
	n := 153
	fmt.Println(isArmstrong(n))
	n = 9474
	fmt.Println(isArmstrong(n))
	n = 9475
	fmt.Println(isArmstrong(n))
}
