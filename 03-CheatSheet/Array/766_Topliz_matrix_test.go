package array

import (
	"fmt"
	"testing"
)

func TestIsToeplizMatrix(t *testing.T){

	// false
	// [[1,2,3,4],[5,1,2,3],[9,5,1,2]] 这个 case 没有过?
	res := isToeplitzMatrix([][]int{{1, 2,3,4},{5, 1,2,3},{9,5, 1,2}})
	t.Log(res)
}

// 没有 AC
// 为什么呢？

// 通过分析 [[1,2,3,4],[5,1,2,3],[9,5,1,2]] 这个 case 没有过?
//

// 时空复杂度都是最佳的差不多
// https://leetcode.com/problems/toeplitz-matrix/solution/
func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix[0]) == 1 {
		return true
	}

	// we notice the number of diagonal is equal to // column of the matrix ,I am wrong
	// the dimenion of the matrix
	// so whatever they give us,
	// we need to find out all the dimenion of this matrix
	// trace of the matrix

	row := len(matrix)
	col := len(matrix[0])

	// fix row
	for i:=0;i < col; i++{
		r := 0
		c := i
		for r+1 < row && c +1 < col {
			if matrix[r][c] != matrix[0][i] {
				fmt.Println(matrix[r][c])
				return false
			}
			r++
			c++
		}
	}

	for i:=0; i < row; i++{
		r := i
		c := 0
		for r+1 < row && c +1 < col {
			if matrix[r][c] != matrix[i][0] {
				fmt.Println(matrix[r][c])
				return false
			}
			r++
			c++
		}
	}

	return true
}


