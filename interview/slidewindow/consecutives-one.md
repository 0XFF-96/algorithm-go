

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
