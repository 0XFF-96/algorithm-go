


- 最长上升子序列（单串问题）

// 简单版本，复杂度是最高的 

```
func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }
    dp := make([]int, len(nums))
    maxLength := 1 // 这个初始化的值很重要

    dp[0] = 1 
    // i 初始化为 1 
    for i:=1; i < len(nums);i ++ {
        dp[i] = 1 
        for j := 0; j < i; j ++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
        maxLength = max(maxLength, dp[i])
    }
    return maxLength
}

func max(i , j int) int {
    if i > j {
        return i 
    } else {
        return j
    }
}
```


- 最长上升子序列的个数 

```
func findNumberOfLIS(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return n 
    }
    length := make([]int, n) // lengths[i] = length of longest ending in nums[i]
    counts := make([]int, n) // count[i] = number of longest ending in nums[i]

    // 一定要初始化 counts 为 1 
    // 为什么
    for i := 0; i < len(counts); i ++ {
        counts[i] = 1 
    }

    for  j := 0; j < n; j ++ {
        for i := 0 ; i < j; i++ {
            if nums[i] < nums[j] {

                // length 和 counts 状态转移的关键点。
                if length[i] >= length[j] {
                    length[j] = length[i] +1

                    // 为什么需要这一行？
                    // 重置序列？
                    // 前序和？
                    counts[j] = counts[i]  
                } else if  length[i] + 1 == length[j] {
                    counts[j] += counts[i]
                }
            }
        }
    }
    longest := 0 
    ans := 0 
    for i :=0 ; i < len(length); i++ {
        longest = max(longest, length[i])
    }

    for i := 0; i < n; i++{
        if length[i] == longest {
            ans += counts[i]
        }
    }

    return ans
}

func max(i , j int) int {
    if i > j {
        return i 
    } else {
        return j
    }
}
```
