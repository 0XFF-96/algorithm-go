package main 

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
    
    if row <= 2 || col <= 2 || ( row == 3 && col == 3){
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
    for r:=0;r < row +3; r++{
        for c:=0; c < col+3; c++{
            if _, found := set[grid[r][c]]; found {
                return false 
            }
            set[grid[r][c]] = true 
        }
    }
    
    leftToRightDiagonal := grid[row][col] + grid[row+1][col+1] + grid[row+2][col+2]
    rightToLeftDiagonal := grid[row+2][col] + grid[row+1][col+1] + grid[row][col+2] 
    
    if leftToRightDiagonal != rightToLeftDiagonal {
        return false 
    }
    return true 
} 