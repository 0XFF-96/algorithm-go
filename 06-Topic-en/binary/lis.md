
# 300 LIS longest increasing subsequence e


```

func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    result := 1
    dp := make([]int, len(nums))
    for i := range dp {
        dp[i] = 1
    }
    
    for i := 0; i < len(dp) - 1; i++ {
        for j := i + 1; j < len(dp); j++ {
            if nums[i] < nums[j] {
                dp[j] = max(dp[i] + 1, dp[j])
                result = max(result, dp[j])
            }
        }
    }
    
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```


```
increasing Triplet Subsequence
Russian Doll Envelopes
Maximum Length of Pair Chain
Number of Longest Increasing Subsequence
Minimum ASCII Delete Sum for Two Strings
Minimum Number of Removals to Make Mountain Array
Find the Longest Valid Obstacle Course at Each Position
```

