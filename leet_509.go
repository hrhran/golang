package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := 0
	b := 1
	temp := 0
	for i := 2; i <= n; i++ {
		temp = a + b
		a = b
		b = temp
	}
	return b
}

func main() {
	n := 10
	fmt.Println(fib(n))
}
