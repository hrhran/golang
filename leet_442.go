package main

import "fmt"

func findDuplicates(nums []int) []int {
    i := 0
    res := make([]int, 0)
    found := make(map[int]bool)
    for ; i < len(nums); i++ {
        if found[nums[i]] {
            res = append(res, nums[i])
        } else {
            found[nums[i]] = true
        }
    }
    return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}
