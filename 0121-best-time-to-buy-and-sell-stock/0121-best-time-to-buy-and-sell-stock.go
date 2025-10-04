func maxProfit(prices []int) int {
    l, r := 0, 1
    maxProfit := 0

    for r < len(prices) {
        if prices[r] > prices[l] {
            profit := prices[r] - prices[l]
            if profit > maxProfit {
                maxProfit = profit
            }
        } else {
            l = r
        }
        r++
    }
    return maxProfit
}
