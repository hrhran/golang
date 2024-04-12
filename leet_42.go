package main

import "fmt"

func trap(height []int) int {
    res, i, max := 0, 0, 0
    ml, mr := make([]int, len(height)), make([]int, len(height))
    for i = 1; i < len(height); i++ {
        if height[i - 1] > max {
            max = height[i - 1]
        }
        ml[i] = max
    }
    max = 0
    for i = len(height) - 2; i >= 0; i-- {
        if height[i + 1] > max {
            max = height[i + 1]
        }
    mr[i] = max
    }
    for i = 0; i < len(height); i++ {
        if ml[i] < mr[i] {
            max = ml[i]
        } else {
            max = mr[i]
        }
        if (max - height[i]) > 0 {
            res += (max - height[i])
        }
    }
    return res
}

func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(height))
}
