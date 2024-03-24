package main

import "fmt"

func findDuplicate(nums []int) int {
    countMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        countMap[nums[i]]++
        if (countMap[nums[i]] > 1){
            return nums[i]
        }
    }
    return -1
}

func main() {
  nums := []int{1, 2, 2, 3}
  fmt.Println(findDuplicate(nums)
}
