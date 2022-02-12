

- max profix 股票的最大利润。 


```

func maxProfit(prices []int) int {
    length := len(prices)
    min := prices[0]
    maxDiff := 0 

    for i := 0; i < length; i ++ {
        maxDiff = max(maxDiff, prices[i] - min)
        if prices[i] < min {
            min = prices[i]
        }
    }
    return maxDiff
}

```

- 股票买卖2。 重点是状态转移和状态压缩的算法
```

func maxProfit(prices []int) int {
    n := len(prices)
    
    // dp[i][0] 表示第i天结束, 手里没有股票的最大收益, dp[i][1]表示第i天结束, 手里有股票的最大收益
    // 进行了状态压缩。
    dp0 := 0 
    dp1 := - prices[0] // 必须先买入。


    for i := 1; i < n; i++{
        // dp1: 昨天持有股票
        // dp0: 昨天没有持有股票

        // dp0 - prices[i], 昨天没有持有股票，所有今天可以买股票。
        // dp1 + prices[i], 昨天持有股票，今天只能卖股票了。
        dp0 = max(dp0, dp1 + prices[i]) //今天不打算持有股票。 前一天没有买入，和今天卖出的最大值。
        dp1 = max(dp1, dp0 - prices[i]) //今天打算持有股票。dp0 - prices[i] 
    }
    return dp0 
}

```
