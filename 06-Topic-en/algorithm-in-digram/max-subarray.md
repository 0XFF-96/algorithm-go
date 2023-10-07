# Max sub array 

- 连续子数组最大和。 

```
func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }
    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0]
    result := dp[0]

    for i := 1; i < n; i++ {
        dp[i] = max(dp[i-1], 0) + nums[i]
        result = max(dp[i], result)
    }
    return result
}
```
