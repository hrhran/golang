package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
    max, i := 0, 0
    res := make([]bool, len(candies))
    for ; i < len(candies); i++ {
      if candies[i] > max {
        max = candies[i]
      }
    }
    for i = 0; i < len(candies); i++ {
      res[i] = candies[i] + extraCandies >= max
    }
    return res
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
	candies = []int{4, 2, 1, 1, 2}
	extraCandies = 1
	fmt.Println(kidsWithCandies(candies, extraCandies))
	candies = []int{12, 1, 12}
	extraCandies = 10
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
