

- 最大子子序和 （与 最大子数组的和）

```

func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }
    // 初始化
    maxSum := nums[0]
    curSum := nums[0]
    for i := 1; i < len(nums); i++ {
        curSum = max(nums[i], curSum + nums[i])
        maxSum = max(maxSum, curSum)
    }
    return maxSum
}

```
