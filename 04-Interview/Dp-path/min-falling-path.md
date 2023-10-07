- 最小下降路径。 

func minFallingPathSum(matrix [][]int) int {
    n := len(matrix)
    // 是一个 matrix ! row = col 
    dp := make([][]int, n) // 初始化不一样
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }

    // 初始化，枚举起点。每个位置的最小成本即本身
    for j := 0;  j < n; j++ {
        dp[0][j] = matrix[0][j]
    }

    for i := 1; i < n; i++ {
        for j := 0; j < n; j++ {

            // 中间情况
            val := matrix[i][j]
            dp[i][j] = dp[i-1][j] + val // 为什么是这行先？// 中间情况。

            // 往左
            if j - 1 >= 0 {
                dp[i][j] = min(dp[i][j], dp[i-1][j-1] + val)
            }

            // 往右
            if j + 1 < n {
                dp[i][j] = min(dp[i][j], dp[i-1][j+1] + val)
            }
        }
    }

    // 遍历最后一行的最小值
    ans := 10000000 
    for i := 0; i < n; i++ {
        ans = min(ans, dp[n-1][i])
    }

    return ans 
}


- 
