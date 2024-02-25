package main

import "fmt"

func isHappy(n int) bool {
	if n == 1 || n == 7 {
		return true
	}
	for n > 9 {
		sum := 0
		digit := 0
		for n > 0 {
			digit = n % 10
			sum = sum + (digit * digit)
			n = n / 10
		}
		if sum == 1 || sum == 7 {
			return true
		}
		n = sum
	}
	return false
}

func main() {
	n := 19
	fmt.Println(isHappy(n))
	n = 2
	fmt.Println(isHappy(n))
	n = 7
	fmt.Println(isHappy(n))
}
