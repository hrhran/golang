package main

import (
	"fmt"
	"strings"
)

func numJewelsInStones(jewels string, stones string) int {
    count := 0
    for i := 0; i < len(stones); i++ {
        if (strings.Contains(jewels, string(stones[i]))) {
            count++
        }
    }
    return count
}

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels, stones))
	jewels = "z"
	stones = "ZZ"
	fmt.Println(numJewelsInStones(jewels, stones))
}
