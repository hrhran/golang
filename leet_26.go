package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 1
	j := 1
	for i < len(nums) {
		if nums[i-1] != nums[i] {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	return j
}

func main() {
	n := []int{1, 1, 2}
	fmt.Println(removeDuplicates(n))
}
