package main

import "fmt"

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    res := 0
    for i := 0; i < len(hours); i++ {
        if (hours[i] >= target) {
            res++
        }
    }
    return res
}


func main() {
    hours := []int{160, 180, 150, 200, 170}
    target := 160
    fmt.Println(numberOfEmployeesWhoMetTarget(hours, target))
}
