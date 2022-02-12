
- 矩阵最小路径和 
- mistake, 一开始写错了状态转移方程。因为括号没有看到的问题。 

```

func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    dp := make([][]int, m)
    for i := 0; i < m; i ++ {
        dp[i] = make([]int, n)
    }

    // 初始化
    dp[0][0] = grid[0][0]
    for i := 1; i < m; i ++ {
        dp[i][0] = dp[i-1][0] + grid[i][0] // 边界条件
    }

    for j := 1; j < n; j++ {
        dp[0][j] = dp[0][j-1] + grid[0][j]
    }

    for i := 1; i < m; i ++ {
        for j := 1; j < n; j ++ {
                //由于[i,j]位置只能由[i - 1, j] 或者[i][j - 1]移动而来
                //状态转移方程：dp[i][j] = min(dp[i - 1][j], dp[i][j - 1])
                // 从左上角到右下角的路径
                // 说明：每次只能向下或者向右移动一步

                // 这里的状态转移很重要‼️
                // 如果是允许， 往下，往右，往右下角移动。这样子的话，这个转移方程就需要改一下了。 
            dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
        }
    }

    return dp[m-1][n-1]
}

```
