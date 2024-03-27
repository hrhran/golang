package main

import "fmt"

func convertTemperature(celsius float64) []float64 {
    return []float64{celsius + 273.15, celsius * 1.80 + 32.00}
}

func main() {
	celsius := 0.0
	fmt.Println(convertTemperature(celsius))
	celsius = 100.0
	fmt.Println(convertTemperature(celsius))
}
