package main

import "fmt"

func strStr(haystack string, needle string) int {
	index := -1
	temp := 0
	i := 0
	j := 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				index = temp
				break
			}
		} else {
			j = 0
			temp++
			i = temp
		}
	}
	return index
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
	haystack = "mississippi"
	needle = "issip"
	fmt.Println(strStr(haystack, needle))
}
