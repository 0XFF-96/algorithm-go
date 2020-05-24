package _4_search_2d_array


// AC
// Runtime: 4 ms, faster than 99.42% of Go online submissions for Search a 2D Matrix.
// Memory Usage: 3.8 MB, less than 100.00% of Go online submissions for Search a 2D Matrix.
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])

	for i := row-1; i >=0; i--{
		if target < matrix[i][0] {
			continue
		} else if target == matrix[i][0] {
			return true
		} else {
			for k:=0; k < col; k++{
				if target == matrix[i][k]{
					return true
				}
			}
			return false
		}
	}
	return false
}
