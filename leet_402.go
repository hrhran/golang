package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	var res []byte
	i := 0
	if len(num) == 0 || len(num) == 1 {
		return "0"
	}
	for i < len(num) {
		if len(res) == 0 || (len(res) > 0 && num[i] >= res[len(res)-1]) || k == 0 {
			res = append(res, num[i])
		} else {
			for len(res) > 0 && num[i] < res[len(res)-1] && k > 0 {
				res = res[:len(res)-1]
				k--
			}
			res = append(res, num[i])
		}
		i++
	}
	for k > 0 {
		res = res[:len(res)-1]
		k--
	}
	result := strings.TrimLeft(string(res), "0")
	if result == "" {
		return "0"
	}
	return result
}

func main() {
	num := "1432219"
	k := 3
	fmt.Println(removeKdigits(num, k))
}
