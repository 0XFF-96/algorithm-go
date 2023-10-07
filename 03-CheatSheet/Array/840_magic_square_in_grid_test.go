package array

import (
	"fmt"
	"testing"
)

//1 <= grid.length <= 10
//1 <= grid[0].length <= 10
//0 <= grid[i][j] <= 15

func TestMagic(t *testing.T){

	// input := [][]int{{1,1,1},{4,5,6},{9,9,9}}
 	// 这个测试用例没有过.. [[4,7,8],[9,5,1],[2,3,6]]  output 1 expected 0
 	//
	input := [][]int{{4,7,8},{9,5,1},{2,3,6}}
	num := numMagicSquaresInside(input)


	is := isMagic(input, 0, 0)
	t.Log(num, is)

}

// limitation:
// 1、distinct numbers.
// 2、[both] diagonals all have the same sum.

// ===
// restriction is
// row >= 3
// col >= 3


// hints
// 1、two  matrix dimansion traverse
// 2、brute force 、二维滑窗、
func numMagicSquaresInside(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	if row <= 2 || col <= 2 {
		return 0
	}
	numOfMagic := 0
	// frame.
	// detail
	for r :=0; r < row; {
		for c :=0; c < col; {
			if isMagic(grid, r, c) {
				numOfMagic++
			}
			c += 3
		}
		r+= 3
	}


	return numOfMagic
}


func isMagic(grid [][]int, row, col int) bool {
	set := make(map[int]bool, (row+1)*(col+1))

	// check unique
	for r:=0;r < row+3; r++{
		for c:=0; c < col+3; c++{
			if _, found := set[grid[r][c]]; found {
				return false
			}
			set[grid[r][c]] = true
		}
	}
	fmt.Println(set)
	leftToRightDiagonal := grid[row][col] + grid[row+1][col+1] + grid[row+2][col+2]
	rightToLeftDiagonal := grid[row+2][col] + grid[row+1][col+1] + grid[row][col+2]

	if leftToRightDiagonal != rightToLeftDiagonal {
		return false
	}
	return true
}
