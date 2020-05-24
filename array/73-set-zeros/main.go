package main



// If any cell of the matrix has a zero
// we can record its row and column number using additional memory.
// But if you don't want to use extra memory then you can manipulate the array instead. i.e.
// simulating exactly what the question says.


// Setting cell values to zero on the fly while iterating might lead to discrepancies.
// What if you use some other integer value as your marker?
// There is still a better approach for this problem with 0(1) space.

// We could have used 2 sets to keep a record of rows/columns
// which need to be set to zero. But for an O(1) space solution,
// you can use one of the rows and and one of the columns to keep track of this information

// We can use the first cell of every row and column as a flag.
// This flag would determine whether a row or column has been set to zero



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


func setZeroesSolution1(matrix [][]int)  {
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
	col := len(matrix[0])

	rowSet := map[int]struct{}{}
	colSet := map[int]struct{}{}


	for i :=0; i < row; i ++{
		for j := 0; j < col; j++{
			if matrix[i][j] == 0 {
				// set row zeros
				// do it in-place
				// matrix[i] = make([]int, )


				rowSet[i] = struct{}{}
				colSet[j] = struct{}{}
				// // fixed row
				//    for k := 0; k < col; k ++ {
				//        matrix[i][k] = 0
				//    }
				// // fixed col
				//    for y := 0; y < row; y++ {
				//        matrix[y][j] = 0
				//    }
			}
		}
	}

	for i := 0; i < row; i ++{
		for j := 0 ; j < col; j++{
			_, ok := rowSet[i]
			_, ok2 := colSet[j]
			if ok || ok2 {
				matrix[i][j] = 0
			}
		}
	}
	return
}


// 这种做法也是错的❌的
// 当同一行 or 列存在 两个 0 的时候，就会不行。
// 例如: [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
func setZeroesWrong2(matrix [][]int)  {
	if len(matrix) == 0 {
		return
	}
	row := len(matrix)
	col := len(matrix[0])

	rowSet := map[int]struct{}{}
	colSet := map[int]struct{}{}


	for i :=0; i < row; i ++{
		for j := 0; j < col; j++{
			if matrix[i][j] == 0 {

				_, ok := rowSet[i]
				if ok {
					continue
				} else {
					rowSet[i] = struct{}{}
				}

				_, ok = colSet[j]
				if ok {
					continue
				} else {
					colSet[j] = struct{}{}
				}

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