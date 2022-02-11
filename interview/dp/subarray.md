

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

- 环形数组的最大和

mistakes 
- 变量初始化和命名不一致
- 漏掉了初始化的关键代码
- 没有能够完全理解代码， maxright 和 rightsum 数组直接有什么关系，为什么需要。
- 不能用代码逻辑解析。 

```

func maxSubarraySumCircular(nums []int) int {
    N := len(nums)

    ans := nums[0]
    cur := nums[0]

    for i:=1; i < N; i++{
        cur = nums[i] + max(cur, 0) // 前缀和？
        ans = max(ans, cur)
    }

    // ans is the answer for 1-interval subarrays
    // Now Let's consider all 2-interval subarrays 

    // For each i, we want to know 
    // the maximum of sum[A[j:]] with j >= i + 2 

    // maxright[i] = max(j >= i) rightsums[j]
    rightsum := make([]int, N)
    rightsum[N-1] = nums[N-1] // 初始化rightsum
    for i := N - 2; i >= 0; i-- {
        // 反向的前缀和
        rightsum[i] = rightsum[i+1] + nums[i] 
    }

    // maxright[i] = max(j >= i) rightsums[j]
    maxright := make([]int, N)
    maxright[N-1] = nums[N-1] // 初始化

    for i := N -2; i >= 0; i--{
        maxright[i] = max(maxright[i+1], rightsum[i])
    }

    leftsum := 0 
    for i := 0; i < N-2; i++{
        leftsum += nums[i]
        ans = max(ans, leftsum + maxright[i+2])
    }
    return ans 
}

```
