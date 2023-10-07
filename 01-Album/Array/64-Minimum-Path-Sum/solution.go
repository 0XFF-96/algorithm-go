package main

import (
	"fmt"
)

// 终于 AC
// 做这条的时候，状态有点差
// 导致基本的条件语句没有写好
// 而且写完之后，没有认真检查一遍
// 导致最终在 debug 上浪费过多时间。
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])

	for i:=0; i < row; i ++{
		for j := 0; j < col; j++{
			fmt.Println("------")
			fmt.Println(grid[i][j])
			if i - 1 < 0 && j - 1 < 0 {
				grid[i][j] = grid[i][j]
			} else if i - 1 >= 0  && j - 1 < 0 {
				grid[i][j] +=  grid[i-1][j]
			} else if  j - 1 >= 0 && i - 1 < 0 {
				grid[i][j] += grid[i][j-1]
			} else {
				grid[i][j] +=  min(grid[i-1][j], grid[i][j-1])
			}
			fmt.Println(grid[i][j])
			fmt.Println("------")
		}
	}
	fmt.Println(grid)
	return grid[row-1][col-1]
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}