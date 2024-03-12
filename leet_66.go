package main

import "fmt"

func plusOne(digits []int) []int {
    sum, carry := 0, 1
    res := make([]int, len(digits) + 1)
    for i := len(digits) - 1; i >= 0; i-- {
        sum = digits[i] + carry
        res[i + 1] = sum % 10
        carry = sum / 10
    }
    if (carry == 1) {
        res[0] = 1
    } else {
        return res[1:]
    }
    return res
}

func main() {
	n := []int{1, 2, 3}
	fmt.Println(plusOne(n))
	n = []int{4, 3, 2, 1}
	fmt.Println(plusOne(n))
}
