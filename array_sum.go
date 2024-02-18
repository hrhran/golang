package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the size of the array: ")
	fmt.Scanln(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scanln(&numbers[i])
	}
	var sum int
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("Sum:", sum)
}
