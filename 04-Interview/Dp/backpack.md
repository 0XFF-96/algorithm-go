

- 零钱兑换问题 （动态规划，二维数组解法）


```
func coinChange(coins []int, amount int) int {
    length := len(coins)

    maxValue := 100000
    dp := make([]int, amount + 1)
    for i := 1; i < len(dp); i ++ {
        dp[i] = maxValue
    }

    dp[0] = 0 
    for i := 1; i <= amount; i++ {
        for j := 0; j < length; j ++ {
            if coins[j] <= i { // 为什么是等于？
                dp[i] = min(dp[i], dp[i - coins[j]] + 1)
            }
        }
    }

    if dp[amount] == maxValue {
        return -1 
    } else {
        return dp[amount]
    }
}
```
