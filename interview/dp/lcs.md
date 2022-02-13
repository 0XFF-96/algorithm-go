


- 最长公共子序列（动态规划问题）
- 需要利用多种不同的方法， 和初始化的方法做一下。

```

func longestCommonSubsequence(text1 string, text2 string) int {
    m := len(text1)
    n := len(text2)

    dp := make([][]int, m + 1)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }

    for i:=0; i < m; i++ {
        dp[i][0] = 0 
    }

    for j:=0; j < n; j++{
        dp[0][j] = 0 
    }

    // 错误❌
    // 为什么 i <= m, j <= n ？？
    for i := 1; i <= m; i++{
        for j := 1; j <= n; j ++ {
            if text1[ i - 1 ] == text2[ j - 1] {
                // 对角线位置 + 1， 相当于删除两元素
                dp[i][j] = dp[ i - 1 ][ j - 1] + 1 
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    fmt.Printf("%+v", dp)
    return dp[m][n]
}

```


- 最长公共子数组


```

func findLength(nums1 []int, nums2 []int) int {
    m := len(nums1)
    n := len(nums2)

    // 已经包含初始化为 0 的操作.
    // 如果需要状态压缩的话， 会从 dp 数组进行入手
    // 常见的解决方案是
    // 1. 滚动数组
    dp := make([][]int, m + 1)
    for i := 0; i < m + 1; i ++ {
        dp[i] = make([]int, n + 1)
    }

    res := 0 
    for i := 1; i < m + 1; i ++ {
        for j := 1; j < n + 1; j ++ {
            if nums1[i-1] == nums2[j-1] {
                 dp[i][j] = dp[i-1][j-1] + 1 // diagonal value + 1 
            } else {
                dp[i][j] = 0 
            }  
             res = max(res, dp[i][j])
        } 
    }
    return res 
}

```
