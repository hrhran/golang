package main

import "fmt"

func romanToInt(s string) int {
	var symbolMap = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && symbolMap[s[i]] < symbolMap[s[i+1]] {
			res -= symbolMap[s[i]]
		} else {
			res += symbolMap[s[i]]
		}
	}
	return res
}

func main() {
	s := "III"
	fmt.Println(romanToInt(s))
	s = "IV"
	fmt.Println(romanToInt(s))
	s = "IX"
	fmt.Println(romanToInt(s))
	s = "LVIII"
	fmt.Println(romanToInt(s))
	s = "MCMXCIV"
	fmt.Println(romanToInt(s))
}
