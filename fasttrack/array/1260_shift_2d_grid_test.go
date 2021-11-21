package array


// 80% 100 %
// 有三种方法做这一个题目。
// 重点!
// 2-d grid:https://leetcode.com/articles/shift-2d-grid/
//
func shiftGrid(grid [][]int, k int) [][]int {
	// 二维坐标，映射到一维坐标里面
	row := len(grid)
	col := len(grid[0])

	tmp := make([]int, row*col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			idx := (i*col + j + k) % (row * col)
			tmp[idx] = grid[i][j]
		}
	}

	for idx := range tmp {
		r := idx / col
		c := idx % col
		grid[r][c] = tmp[idx]
	}
	// two dimension array to one dimension ?
	// how to reflection it?
	return grid
}



