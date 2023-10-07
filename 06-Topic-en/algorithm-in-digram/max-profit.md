# 股票最大利润

```

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0 
    } 
    dp := make([]int, len(prices))
    
    cost := prices[0]
    for i := 1; i < len(prices); i ++ {
        cost = min(cost, prices[i])
        dp[i] = max(dp[i-1], prices[i] - cost)
    }
    return dp[len(prices) -1]
}

```
