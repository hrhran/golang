package main

import "fmt"

func defangIPaddr(address string) string {
    result := make([]byte, 0, len(address)+6)
    for i := 0; i < len(address); i++ {
        if address[i] == '.' {
            result = append(result, '[', '.', ']')
        } else {
            result = append(result, address[i])
        }
    }
    return string(result)
}


func main() {
	address := "1.1.1.1"
	fmt.Println(defangIPaddr(address))
	address = "255.100.50.0"
	fmt.Println(defangIPaddr(address))
}