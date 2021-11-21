package array

func transpose(A [][]int) [][]int {
	// 二分
	// matrix 行列相等？如果不是，没办法 traspose
	//  Rows become columns and vice versa!
	//
	row := len(A)
	col := len(A[0])
	// init
	ret := make([][]int, col)
	for i:=0;i < col;i++{
		ret[i] = make([]int, row)
	}
	// 可以减少一半循环的方法？
	for r:=0;r <row;r++{
		for c:=0;c<col;c++{
			ret[c][r] = A[r][c]
		}
	}
	return ret
}

func transposeV2(A [][]int) [][]int {
	if len(A) == 0{
		return nil
	}
	row := len(A)
	col := len(A[0])

	ret := make([][]int, col)
	for i:=0;i < col;i++{
		ret[i] = make([]int, row)
	}

	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			ret[j][i] = A[i][j]
		}
	}
	return ret
}
