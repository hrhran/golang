package main

import "fmt"

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    i, res := 0, 0
    for ; i < len(items); i++ {
        if ruleKey == "type" {
            if (items[i][0] == ruleValue) {
                res++
            }
        } else if ruleKey == "color" {
            if (items[i][1] == ruleValue) {
                res++
            }
        } else {
            if (items[i][2] == ruleValue) {
                res++
            }
        }
    }
    return res
}

func main() {
	items := [][]string{{"phone","blue","pixel"},{"computer","silver","lenovo"},{"phone","gold","iphone"}}
	ruleKey := "color"
	ruleValue := "silver"
	fmt.Println(countMatches(items, ruleKey, ruleValue))
	ruleKey = "type"
	ruleValue = "phone"
	fmt.Println(countMatches(items, ruleKey, ruleValue))
}
