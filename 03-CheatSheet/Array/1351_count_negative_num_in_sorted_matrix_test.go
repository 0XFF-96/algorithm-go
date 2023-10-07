package array



func countNegatives(grid [][]int) int {
	// 从最右开始起
	count := 0
	for i:=0; i<len(grid); i++{
		for j:=0; j < len(grid[i]);j++{
			if grid[i][j] < 0 {
				count++
			}
		}
	}
	return count
}

// 一定存在更快的方法..
//
