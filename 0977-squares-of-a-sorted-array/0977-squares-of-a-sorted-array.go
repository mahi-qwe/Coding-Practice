import (
    "sort"
    // "math"
)

func sortedSquares(nums []int) []int {
    sqr := []int{}
    for _, num := range nums{
        sqr = append(sqr, num*num)
    }
    sort.Ints(sqr)
    return sqr
}