package main

import "fmt"

func isVowel(c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func reverseVowels(s string) string {
    res := []byte(s)
    i, vowels := 0, ""
    for i < len(res) {
        if (isVowel(res[i])) {
        vowels = fmt.Sprintf("%s%c", vowels, res[i])
        }
        i++
    }
    j := len(vowels) - 1
    for i = 0; i < len(res); i++ {
        if (isVowel(res[i])) {
            res[i] = vowels[j]
            j--
        }
    }
    return string(res)
}

func main() {
	s := "hello"
	fmt.Println(reverseVowels(s))
	s = "leetcode"
	fmt.Println(reverseVowels(s))
	s = "aA"
	fmt.Println(reverseVowels(s))
}
