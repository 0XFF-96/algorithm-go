package array

import (
	"math"
)

// solution: 应该有更好的方法解题
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solution/
//
// The points of interest are the peaks and
// valleys in the given graph. We need to find
// the largest peak following the smallest valley.
// We can maintain two variables
// 效率太低了...
func maxProfit(prices []int) int {
	max:=0
	for idx, v := range prices {
		for i:=idx; i < len(prices); i++{
			tmp := prices[i] -v
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}


// 空间利用率，还可以怎么提升呢？
func maxProfitV2(prices []int) int {
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