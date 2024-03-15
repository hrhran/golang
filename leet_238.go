package main

import "fmt"

func productExceptSelf(nums []int) []int {
    i, product, zero, flag := 0, 1, 0, false
    res := make([]int, len(nums))
    for i = 0; i < len(nums); i++ {
        if nums[i] != 0 {
            product *= nums[i]
        } else {
            zero++
            flag = true
        }
    }
    if (zero > 1) {
        return res
    }
    for i = 0; i < len(nums); i++ {
        if flag == true {
            if nums[i] != 0 {
                res[i] = 0
            } else {
                res[i] = product
            }
        } else {
            res[i] = product/nums[i]
        }
    }
    return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
	nums = []int{0, 0}
	fmt.Println(productExceptSelf(nums))
}
