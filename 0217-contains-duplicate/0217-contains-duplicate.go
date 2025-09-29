func containsDuplicate(nums []int) bool {
    seenMap := make(map[int]bool)
    for _, num := range nums{
        if seenMap[num] {
            return true
        }
        seenMap[num] = true
    }
    return false
}