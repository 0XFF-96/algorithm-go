package main

// 错误的做法
func setZeroes(matrix [][]int)  {
	if len(matrix) == 0 {
		return
	}
	// 难点在于：1、设置过一次的行和列，不需要被重复设置
	// 2、当一行/一列，存在重复设置过的值时，该怎么办？


	// 有一点没有考虑
	// 有些，因为因为之前遍历的时候，就被设置过了 0
	// 之后，就不再需要被设置了...

	// 所有，需要记录哪些行,列被曾经被设置过？
	// Could you devise a constant space solution?

	// 在设置的时候，检查一下啊，这里的 col, index 是否已经被设置过 0
	// 然后，这里的参数是否为 0
	row := len(matrix)
	col := len(matrix) -1

	for i :=0; i < row; i ++{
		for j := 0; j < col; j++{
			if matrix[i][j] == 0 {
				// set row zeros
				// do it in-place
				// matrix[i] = make([]int, )

				// fixed row
				for k := 0; k < col; k ++ {
					matrix[i][k] = 0
				}
				// fixed col
				for y := 0; y < row; y++ {
					matrix[y][j] = 0
				}
			}
		}
	}

	return
}

