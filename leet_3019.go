package main

import "fmt"

func countKeyChanges(s string) int {
    count := 0
    for i := 1; i < len(s); i++ {
        if !((int(s[i - 1]) - int(s[i])) == 0 || (int(s[i - 1]) - int(s[i])) == 32 || (int(s[i - 1]) - int(s[i])) == -32) {
            count++
        }
    }
    return count
}

func main() {
	s := "aAbc"
	fmt.Println(countKeyChanges(s))
	s = "bfaBb"
	fmt.Println(countKeyChanges(s))
}
