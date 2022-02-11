


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
