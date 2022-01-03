package main 


// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfitV1(prices []int) int {
    // not engaing multiple transactions at the same time
    // 只要后面有收益，就 buy
    // 简单的二维搜索问题
    // for-for 循环，暴力点就
    // [7, 1, 5, 3, 6, 4]
    // [1, 2, 3, 4, 5]
    // Time complexity : O(n^n)O(n 
    // n
    // ). Recursive function is called n^2
    // Space complexity : O(n)O(n). Depth of recursion is n

}

// solution 2 
func maxProfitV2(prices []int) int {
	profit, sell := 0, false

	for i := 0; i < len(prices); i++ {
		if sell {
			profit, sell = profit+prices[i]-prices[i-1], false
		}
		sell = i < len(prices)-1 && prices[i] < prices[i+1]
	}
	return profit
}