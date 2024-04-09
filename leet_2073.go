package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
    res := 0
	for tickets[k] != 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] != 0 && tickets[k] != 0 {
				res++
                tickets[i]--
			}
		}
	}
	return res
}

func main() {
	tickets := []int{2,6,3,4,5}
	k := 2
	fmt.Println(timeRequiredToBuy(tickets, k))
}
