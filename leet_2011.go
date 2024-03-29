package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
    res, i := 0, 0
    for ; i < len(operations); i++ {
        if operations[i][1] == '+' {
            res++
        }
        if operations[i][1] == '-' {
            res--
        }
    }
    return res
}

func main() {
	operations := []string{"--X", "X++", "X++"}
	fmt.Println(finalValueAfterOperations(operations))
}
