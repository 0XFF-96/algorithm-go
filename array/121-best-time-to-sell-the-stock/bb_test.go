

// If you were only permitted to complete at most one transaction 
// (i.e., buy one and sell one share of the stock),

import "math"
func maxProfit(prices []int) int {
    minPrice := math.MaxInt32
    maxProfit := 0 
    
    for i:=0;i < len(prices);i++{
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else if prices[i] - minPrice > maxProfit {
            maxProfit = prices[i] - minPrice
        }
    }
    return maxProfit
}



https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/solution/

