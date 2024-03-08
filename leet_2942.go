package main

import "fmt"

func findWordsContaining(words []string, x byte) []int {
    res, i, j := []int{}, 0, 0
    for i = 0; i < len(words); i++ {
        j = 0;
        for j < len(words[i]) {
            if (x == words[i][j]) {
                res = append(res, i)
                break
            }
            j++
        }
    }
    return res
}

func main() {
	words := []string{"hello", "world", "leetcode"}
	x := byte('e')
	fmt.Println(findWordsContaining(words, x))
}
