package main

import "fmt"

func findMaxLength(nums []int) int {
    max, count, i := 0, 0, 0
    dir := map[int]int {
        0: -1,
    }
    for i = 0; i < len(nums); i++ {
        if nums[i] == 0 {
            count--
        } else {
            count++
        }
        if _, ok := dir[count]; ok {
            if max < i - dir[count] {
                max = i - dir[count]
            }
        } else {
            dir[count] = i
        }
    }
    return max
}

func main() {
	nums := []int{0, 1}
	fmt.Println(findMaxLength(nums))
	nums = []int{0, 1, 0}
	fmt.Println(findMaxLength(nums))
	nums = []int{0, 0, 1, 0, 0, 0, 1, 1}
	fmt.Println(findMaxLength(nums))
}
