
- string match 最小编辑距离。

```

func minDistance(word1 string, word2 string) int {
    // dp[i][j] =
    // 1. dp[i][j-1] + 1  (最后一步是插入)
    // 2. dp[i-1][j] + 1  (最后一步是删)
    // 3. dp[i-1][j-1] + 1  (最后一步是改，且 word1[i] != word2[j])
    // 4. dp[i-1][j-1]  (最后一步是改，且 word1[i] == word2[j])

    // dp 数组的初始化（重点）
    m := len(word1)
    n := len(word2)

    dp := make([][]int, m + 1)
    for i:=0; i < m +1; i++{
        dp[i] = make([]int, n + 1)
    }

    // 真正的初始化值
    for i:=0; i < m + 1 ; i++{
        dp[i][0] = i 
    }

    for j :=0; j < n + 1; j++{
        dp[0][j] = j 
    }
    // 如何遍历顺序。
    // 为什么从 1 开始。边界条件的处理
    for i := 1; i < m + 1; i ++ {
        for j := 1; j < n + 1 ; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                delete := dp[i-1][j] + 1 
                insert := dp[i][j-1] + 1 
                update := dp[i-1][j-1] + 1 
                dp[i][j] =min(min(delete, insert), update)
            }
        }
    }

    return dp[m][n]
}

```
