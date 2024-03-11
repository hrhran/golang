package main

import "fmt"

func smallestEvenMultiple(n int) int {
    if n % 2 == 0 {
        return n
    }
    return n * 2
}

func main() {
	n := 3
	fmt.Println(smallestEvenMultiple(n))
}
