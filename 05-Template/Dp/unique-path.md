### 不同路径问题

- 包含障碍物的不同路径，主要影响如何初始化二维数组和计算 dp 数据的最大值。

### 不同路径一

```
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
```




### 不同路径二

```
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m,n:= len(obstacleGrid),len(obstacleGrid[0])
	// 定义一个dp数组
	dp := make([][]int,m)
	for i,_ := range dp {
		dp[i] = make([]int,n)
	}
	// 初始化
	for i:=0;i<m;i++ {
		// 如果是障碍物, 后面的就都是0, 不用循环了
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0]=1
	}
	for i:=0;i<n;i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i]=1
	}
	// dp数组推导过程
	for i:=1;i<m;i++ {
		for j:=1;j<n;j++ {
			// 如果obstacleGrid[i][j]这个点是障碍物, 那么我们的dp[i][j]保持为0
           	 	if obstacleGrid[i][j] != 1 {
				// 否则我们需要计算当前点可以到达的路径数
				dp[i][j] = dp[i-1][j]+dp[i][j-1]
			}
		}
	}
	// debug遍历dp
	//for i,_ := range dp {
	//	for j,_ := range dp[i] {
	//		fmt.Printf("%.2v,",dp[i][j])
	//	}
	//	fmt.Println()
	//}
	return dp[m-1][n-1]
}
```