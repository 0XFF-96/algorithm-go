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


