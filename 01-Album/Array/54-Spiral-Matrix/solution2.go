// 按照上下左右方向感递进
func spiralOrder(matrix [][]int) []int {
    var res []int 
    if len(matrix) == 0 {
        return res
    }
    
    rowBegin := 0 
    rowEnd := len(matrix) -1
    colBegin := 0 
    colEnd := len(matrix[0])  -1 
    
    for rowBegin <= rowEnd && colBegin <= colEnd {
        // traverse Right 
        
        for i:=colBegin; i <= colEnd; i ++{
            res = append(res, matrix[rowBegin][i])
        }
        rowBegin++
        
        // traver down 
        for j := rowBegin; j <= rowEnd; j ++{
            res = append(res, matrix[j][colEnd])
        }
        colEnd--
        
        if (rowBegin <= rowEnd) {
            // Traverse Left 
            for j := colEnd; j >= colBegin; j-- {
                res = append(res, matrix[rowEnd][j])
            }
        }
        rowEnd--
        
        if colBegin <= colEnd {
            // Traver Up 
            for j := rowEnd;j >= rowBegin; j --{
                res = append(res, matrix[j][colBegin])
            }
        }
        colBegin++
        
    }
    

    return res
}