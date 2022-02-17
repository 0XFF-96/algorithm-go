
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

- 无重复字符的最长子串。
- 应该还有其他很多解法，需要整理一下。

```
func lengthOfLongestSubstring(s string) int {
    m := map[byte]int{}
    n := len(s)

    rk, ans := -1, 0 // 为什么初始化为这个？

    // 这道题目一定需要理解窗口的收缩方向是什么？
    // 窗口是如何进行收缩的？
    for i := 0; i < n; i ++ {
        if i != 0 {
            delete(m, s[i-1]) // 左指针移动，删除一个字符
        }
        for rk + 1 < n && m[s[rk + 1]] == 0 {
            // 不断移动右指针
            m[s[rk+1]] += 1
            rk++ 
        }

        // 第 i 到 rk 个字符是一个最长的无重复字符串。 
        ans = max(ans, rk - i + 1) // 为什么需要 减1 ？
    }
    return ans 
}
```
