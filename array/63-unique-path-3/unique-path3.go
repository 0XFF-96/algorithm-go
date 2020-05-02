package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 初始化矩阵 ...
	// 相关行列 ...
	//
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	obstacleGrid[0][0] = 1
	// [[1]] [[0]]
	// 这种情况怎么办？
	//
	// [[0,0],[1,1],[0,0]]
	// 这种情况也不过

	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	for i:=1; i < row; i ++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}

	for i:=1; i < col; i ++ {
		if obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1 ] == 1 {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
		}
	}

	for i := 1; i < row; i++{
		for j:=1; j < col; j++{
			cur := 0
			if j - 1 >= 0 {
				cur += obstacleGrid[i][j-1]
			}
			if i - 1 >= 0 {
				cur += obstacleGrid[i-1][j]
			}

			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = cur
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}

	// fmt.Println(obstacleGrid)
	return obstacleGrid[row-1][col-1]
}
