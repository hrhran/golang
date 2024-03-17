package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
    i, temp := 0, 0
    var res [][]int
    for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
        res = append(res, intervals[i])
    }
    for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
        if (newInterval[0] < intervals[i][0]) {
            temp = newInterval[0]
        } else {
            temp = intervals[i][0]
        }
        newInterval[0] = temp
        temp = 0
        if (newInterval[1] > intervals[i][1]) {
            temp = newInterval[1]
        } else {
            temp = intervals[i][1]
        }
        newInterval[1] = temp
    }
    res = append(res, newInterval)
    for ; i < len(intervals); i++ {
        res = append(res, intervals[i])
    }
    return res
}

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	fmt.Println(insert(intervals, newInterval))
	intervals = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval = []int{4, 8}
	fmt.Println(insert(intervals, newInterval))
	intervals = [][]int{}
	newInterval = []int{5, 7}
	fmt.Println(insert(intervals, newInterval))
	intervals = [][]int{{1, 5}}
	newInterval = []int{2, 3}
	fmt.Println(insert(intervals, newInterval))
}
