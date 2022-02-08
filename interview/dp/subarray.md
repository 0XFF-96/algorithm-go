

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


- 乘积最大的子数组

```
func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }

    // 初始化
    maxF := nums[0]
    minF := nums[0]
    ans := nums[0]

    for i := 1; i < len(nums); i++{
        mx := maxF 
        mn := minF 
        maxF = max(maxF * nums[i], max(nums[i], minF * nums[i]))
        minF = min(mn * nums[i], min(nums[i], mx * nums[i]))

        ans = max(maxF, ans)
    }

    return ans 
}
```

