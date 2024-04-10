package main

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
    var que []int
    res, i := make([]int, len(deck)), 0
    sort.Ints(deck)
    for ; i < len(deck); i++ {
        que = append(que, i)
    }
    for i = 0; i < len(deck); i++ {
        res[que[0]] = deck[i]
        que = que[1:]
        if len(que) == 0 {
            break
        }
        que = append(que[1:], que[0])
    }
    return res
}

func main() {
	n := []int{17,13,11,2,3,5,7}
	fmt.Println(deckRevealedIncreasing(n))
}
