- counting , 
```

func numSubarrayBoundedMax(nums []int, left int, right int) int {
    // 这里为啥不用加 1， 而是需要 left - 1 
    return lessEqualsThan(nums, right) - lessEqualsThan(nums, left - 1)
}

func lessEqualsThan(nums []int, k int) int {
    length := len(nums)
    res := 0 
    left, right := 0, 0 
    for right < length {
        if nums[right] > k {
            left = right + 1 
        }
        res += right - left 
        right++
    }
    return res 
}

```
