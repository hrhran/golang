package main

import "fmt"

func decToBin(n int) string {
	res := ""
	for n != 0 {
		val := n % 2
		res = fmt.Sprintf("%v%v", val, res)
		n = n / 2;
	}
	return res
}

func main() {
	n := 10
	fmt.Println(decToBin(n))
	n = 19
	fmt.Println(decToBin(n))
	n = 7
	fmt.Println(decToBin(n))
}
