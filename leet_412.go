package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	arr := make([]string, n, n)
	i := 0
	for i < n {
		switch {
		case (((i+1)%3) == 0 && ((i+1)%5 == 0)):
			arr[i] = "FizzBuzz"
		case (((i + 1) % 3) == 0):
			arr[i] = "Fizz"
		case (((i + 1) % 5) == 0):
			arr[i] = "Buzz"
		default:
			arr[i] = strconv.Itoa(i + 1)
		}
		i++
	}
	return arr
}

func main() {
	n := 15
	fmt.Println(fizzBuzz(n))
}
