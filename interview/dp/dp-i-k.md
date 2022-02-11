
- 最大平均值和的分组
- 不会， 进制转换的时候出现了问题。 

```

func largestSumOfAverages(nums []int, k int) float64 {
    length := len(nums)

    // 初始化二维数组的初始化
    dp := make([][]float64, length + 1)
    for i, _ := range dp {
        dp[i] = make([]float64, k + 1)
    }

    // 前缀和的计算
    prefixsum := make([]int, length + 1)
    for i := 1; i <= length; i++ {
        prefixsum[i] = prefixsum[i - 1] + nums[i - 1]
    }

    // 再次初始化 dp 的初始值
    for i := 1; i <= length; i ++ {
        dp[i][1] = float64(prefixsum[i]) / float64(i) 
    }

    // 从 2 开始
    for ik := 2; ik <= k; ik ++ {
        // i 从k 开始 
        // i 是小于等于。。。可以等于？
        for i := ik; i <= length; i++ {
            // 状态转移
            // dp[i][k] = 
            // max(dp[j][ik - 1] + avg(j+1, i), dp[i][k])

            for j := ik - 1; j + 1 <= i; j ++ {
                // 计算 j + 1 的和
                avg := float64(prefixsum[i] - prefixsum[j]) / float64(i - (j+1) + 1)
                dp[i][ik] = max(dp[i][ik], dp[j][ik - 1] + avg)
            }
        }
    }

    return float64(dp[length][k])
}
```
