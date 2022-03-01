- 不同的路径和

```

func uniquePaths(m int, n int) int {
    f := make([][]int, m)
    for i, _ := range f {
        f[i] = make([]int, n)
    }
    f[0][0] = 1 // 初始化值 

    for i := 0; i < m; i++ {
        for j :=0; j < n; j++ {
            if (i > 0 && j > 0) {
                f[i][j] = f[i-1][j] + f[i][j-1] // 转移
            } else if i > 0 {
                f[i][j] = f[i-1][j]
            } else if j > 0 {
                f[i][j] = f[i][j-1]
            }
        }
    }
    return f[m-1][n-1]
}

```
- 有障碍的 Uni-path 。 
```

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])

    // 初始化
    dp := make([][]int, m)
    for idx, _ := range dp {
        dp[idx] = make([]int, n)
    }

    // 初始化的时候，不一样。
    if obstacleGrid[0][0] == 1 {
        dp[0][0] = 0 
    } else {
        dp[0][0] = 1 
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j ++ {
            if obstacleGrid[i][j] != 1 {
                if i > 0 && j > 0 {
                    dp[i][j] = dp[i-1][j] + dp[i][j-1]
                } else if i > 0 {
                    dp[i][j] = dp[i-1][j]
                } else if j > 0 {
                    dp[i][j] = dp[i][j-1]
                }
            }
        }
    }

    // 如果是最后有 障碍物，是不是都不能到达了哇。 
    return dp[m-1][n-1]
}

```


```

class Solution {
    public int minimumTotal(List<List<Integer>> tri) {
        int n = tri.size();
        int ans = Integer.MAX_VALUE;
        int[][] f = new int[n][n];
        f[0][0] = tri.get(0).get(0);
        for (int i = 1; i < n; i++) {
            for (int j = 0; j < i + 1; j++) {
                int val = tri.get(i).get(j);
                f[i][j] = Integer.MAX_VALUE;
                if (j != 0) f[i][j] = Math.min(f[i][j], f[i - 1][j - 1] + val);
                if (j != i) f[i][j] = Math.min(f[i][j], f[i - 1][j] + val);
            }
        }
        for (int i = 0; i < n; i++) ans = Math.min(ans, f[n - 1][i]);
        return ans;
    }
}

```
