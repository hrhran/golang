package main

import "fmt"

func islandPerimeter(grid [][]int) int {
    isle, adj := 0, 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                if j - 1 >= 0 && grid[i][j - 1] == 1 {
                    adj++
                }
                if i - 1 >= 0 && grid[i - 1][j] == 1 {
                    adj++
                }
                isle++
            }
        }
    }
    return (isle * 4) - (adj * 2)
}

func main() {
	grid := [][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}}
	fmt.Println(islandPerimeter(grid))
}
