package main

import (
	"fmt"
	"strings"
)

func minRemoveToMakeValid(s string) string {
    res := []byte(s)
    stack := []int{}
    i, n := 0, 0
    for ; i < len(s); i++ {
        if (s[i] == '(') {
            stack = append(stack, i)
            n++
        }
        if (s[i] == ')') {
            if (len(stack) > 0) {
                stack = stack[:len(stack) - 1]
                n--
            } else {
                res[i] = '#'
            }
        }
    }
    for i = 0; i < n; i++ {
        res[stack[i]] = '#'
    }
    final := strings.Replace(string(res), "#", "", -1)
    return final
}

func main() {
	s := "lee(t(c)o)de)"
	fmt.Println(minRemoveToMakeValid(s))
}
