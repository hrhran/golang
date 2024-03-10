package main

import "fmt"

func searchInsert(nums []int, target int) int {
    if target <= nums[0] {
        return 0
    }
    if target > nums[len(nums) - 1] {
        return len(nums)
    }
    for i := 1; i < len(nums); i++ {
        if target <= nums[i] {
            return i
        }
    }
    return 0
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))
	fmt.Println(searchInsert(nums, 1))
}
