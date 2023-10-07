

// longestOnes 。
// 最大连续 1 的个数。

```

func findMaxConsecutiveOnes(nums []int) int {
    // 进阶：如果输入的数字是作为 无限流 逐个输入如何处理？换句话说，内存不能存储下所有从流中输入的数字。您可以有效地解决吗？

    left := 0 
    right := 0 
    maxCount := 0
    freq := 0 
    
    for right < len(nums) {
        if (nums[right] == 1 ) {
            freq++
        }
        maxCount = max(maxCount, freq)
        // 窗口往右移动
        right++

        // 窗口收缩
        if right - left - maxCount > 1 {
            if nums[left] == 1 {
                freq--
            }
            left++
        }
    }
    return right - left 
}

```


- (不定窗口滑动字符串） 
- 删除一个元素后，最长为 1 的子数组长度。 

```

func longestSubarray(nums []int) int {
    // 三指针
    res := 0 
    cnt := 0 
    left := 0 
    for right := 0; right < len(nums); right++ {
        if nums[right] == 0 {
            cnt++
        }
        for cnt > 1 { // 从中删掉一个元素
            if nums[left] == 0 { // 移动指针了
                cnt-- // 为什么是遇到 0 才减下来？
            }
            left++
        }
        res = max(res, right - left) // 为什么不需要加 1 。
    }
    return res 
}

```
