package main

import "fmt"

func maxSubArraySum(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, []int{}
	}
	maxSum := nums[0]
	currentSum := nums[0]
	start, end := 0, 0
	tempStart := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > currentSum+nums[i] {
			currentSum = nums[i]
			tempStart = i
		} else {
			currentSum += nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
			start = tempStart
			end = i
		}
	}
	return maxSum, nums[start : end+1]
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum, maxSubarray := maxSubArraySum(nums)
	fmt.Println("Maximum Subarray Sum:", maxSum)
	fmt.Println("Maximum Subarray:", maxSubarray)
}
