package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
    if len(points) == 0 {
        return 0
    }
    arrows := 0
    sort.Slice(points, func(i int, j int) bool {
        return points[i][1] < points[j][1]
    })
    last := points[0][1]
    for i := 1; i < len(points); i++ {
        if points[i][0] > last {
            arrows++
            last = points[i][1]
        } else {
            if points[i][1] < last {
                last = points[i][1]
            }
        }
    }
    return arrows + 1
}

func main() {
	n := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(n))
}
