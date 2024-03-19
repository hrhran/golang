package main

import "fmt"

func sum(num1 int, num2 int) int {
    return num1 + num2
}

func main() {
	num1 := 1
	num2 := 2
	fmt.Println(sum(num1, num2))
	num1 = 2
	num2 = 3
	fmt.Println(sum(num1, num2))
}
