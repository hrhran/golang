
 package main
 
 import "fmt"
 
 func numIdenticalPairs(nums []int) int {
     i, j, count := 0, 0, 0
     for ; i < len(nums); i++ {
         for j = i + 1; j < len(nums); j++ {
             if nums[i] == nums[j] {
                 count++
             }
         }
     }
     return count
 }
 
 func main() {
     n := []int{1, 2, 3, 1, 1, 3}
     fmt.Println(numIdenticalPairs(n))
 }
