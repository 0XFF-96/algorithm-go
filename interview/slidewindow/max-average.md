
- 子数组最大平均数 I

```
func findMaxAverage(nums []int, k int) float64 {
    // if len(nums) < k {
    //     return -1 
    // }

    // 1. 提前求出第一个窗口的和。
    windowSum := 0 
    for i := 0; i < k; i ++ {
        windowSum += nums[i]
    }

    // 遍历
    res := windowSum 
    for right := k; right < len(nums); right++ {
        // 此时 nums[right], 未被加入 nums 
        // 初始化时的区间是 [right -k, right) 
        windowSum = windowSum + nums[right] - nums[right - k]
        res = max(res, windowSum)
    }
    return float64(res)/ float64(k) 
}
```


- 盛水最多的容器

```

func maxArea(height []int) int {
    maxArea := -1 
    left := 0 
    right := len(height) - 1 // 能够取到的 
    for left < right {
        maxArea = max(maxArea, min(height[left], height[right]) * (right - left))

        // 木桶原理，短板效应
        if height[left] < height[right] { // 哪个边短，向哪边移动。
            left++
        } else {
            right--
        }
    }
    return maxArea
}

```
