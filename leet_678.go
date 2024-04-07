package main

import "fmt"

func checkValidString(s string) bool {
    i, ob, o := 0, 0, 0
    if (s[0] == ')' || s[len(s) - 1] == '(') {
        return false;
    }
    for ; i < len(s); i++ {
        if (s[i] == '(') {
            ob++
            o++
        } else if (s[i] == ')') {
            ob--
            o--
        } else {
            ob--
            o++
        }
        if o < 0 {
            return false
        }
        if ob < 0 {
            ob = 0
        }
    }
    return ob == 0
}

func main() {
	s := "(*))"
	fmt.Println(checkValidString(s))
}
