package main

func setZeroesV2(matrix [][]int) {
	mlen := len(matrix)
	if mlen == 0 {
		return
	}
	nlen := len(matrix[0])

	zeroRow, zeroCol := false, false
	for m := 0; m < mlen; m++ {
		for n := 0; n < nlen; n++ {
			mark(matrix, m, n, &zeroRow, &zeroCol)
		}
	}
	for m := mlen - 1; m >= 0; m-- {
		for n := nlen - 1; n >= 0; n-- {
			set(matrix, m, n, zeroRow, zeroCol)
		}
	}
}

func mark(matrix [][]int, m, n int, zeroRow, zeroCol *bool) {
	switch {
	case m == 0 && n == 0:
		if matrix[m][n] == 0 {
			*zeroRow, *zeroCol = true, true
		}
	case m == 0:
		if matrix[m][n] == 0 {
			*zeroRow = true
		}
	case n == 0:
		if matrix[m][n] == 0 {
			*zeroCol = true
		}
	default:
		if matrix[m][n] == 0 {
			matrix[m][0] = 0
			matrix[0][n] = 0
		}
	}
}

func set(matrix [][]int, m, n int, zeroRow, zeroCol bool) {
	switch {
	case m == 0 && n == 0:
		if zeroRow || zeroCol {
			matrix[m][n] = 0
		}
	case m == 0:
		if zeroRow {
			matrix[m][n] = 0
		}
	case n == 0:
		if zeroCol {
			matrix[m][n] = 0
		}
	default:
		if matrix[0][n] == 0 || matrix[m][0] == 0 {
			matrix[m][n] = 0
		}
	}
}
