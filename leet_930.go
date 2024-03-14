package main

import "fmt"

func numSubarraysWithSum(nums []int, goal int) int {
    res, sum, i, j := 0, 0 , 0, 0
    for i = 0; i < len(nums); i++ {
        j = i
        sum = 0
        for sum <= goal && j < len(nums) {
            sum = sum + nums[j]
            if sum == goal {
                res = res + 1
            }
            j++
        }
    }
    return res
}

func main() {
	nums := []int{1, 0, 1, 0, 1}
	goal := 2
	fmt.Println(numSubarraysWithSum(nums, goal))
	nums = []int{0, 0, 0, 0, 0}
	goal = 0
	fmt.Println(numSubarraysWithSum(nums, goal))
}
