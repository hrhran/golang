package main

import "fmt"

func shuffle(nums []int, n int) []int {
    i := 0
    res := make([]int, 0, n * 2)
    for i < n {
        res = append(res, nums[i])
        res = append(res, nums[i + n])
        i++
    }
    return res
}

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	fmt.Println(shuffle(nums, 3))
}
