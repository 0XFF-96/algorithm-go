
- 简单滑动窗口值。
```

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }

    low, fast := 0, 0 
    for fast < len(nums) {
        if nums[low] != nums[fast] {
            low += 1
            nums[low] = nums[fast]
        }
        fast += 1 
    }
    return low + 1 
}

```

- 最长连续递增序列


```

func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }

    ret := 1
    for i := 0; i < len(nums); i ++ {
        cur := 1 
        for j := i + 1; j < len(nums); j ++ {
            if nums[j-1] < nums[j] {
                cur += 1 
            } else {
                break 
            }
        }
        ret = max(ret, cur)
    }
    return ret 
}

```
