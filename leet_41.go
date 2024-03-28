func firstMissingPositive(nums []int) int {
    n := len(nums)
    dir := make(map[int]bool)
    for _, i := range nums {
        if i > 0 {
            dir[i] = true
        }
    }

    for i := 1; i < n + 1; i++ {
        if _, found := dir[i]; !found {
            return i
        }
    }
    return n + 1
}
