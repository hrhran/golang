package main

import "fmt"

func lengthOfLastWord(s string) int {
    size := 0
    for i := len(s) - 1; i >= 0; i-- {
        if (s[i] != byte(' ')) {
            size += 1
        } else {
            if size > 0 {
                return size
            }
        }
    }
    return size
}

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
	s = "a "
	fmt.Println(lengthOfLastWord(s))
}
