func generateMatrix(n int) [][]int {
    if n <= 0 {
        return nil 
    } 

    // 构造矩阵
    matrix := make([][]int, n)
    for i := 0; i < n; i ++{
        matrix[i] = make([]int, n)
    }
    
    rowBegin := 0 
    rowEnd := n - 1
    colBegin := 0 
    colEnd := n - 1
    
    count := 1
    for rowBegin <= rowEnd && colBegin <= colEnd {
        // traverse Right 
        
        for i:=colBegin; i <= colEnd; i ++{
            // res = append(res, matrix[rowBegin][i])
            matrix[rowBegin][i] = count 
            count++
        }
        rowBegin++
        
        // traver down 
        for j := rowBegin; j <= rowEnd; j ++{
            // res = append(res, matrix[j][colEnd])
            matrix[j][colEnd] = count 
            count++
        }
        colEnd--
        
        if (rowBegin <= rowEnd) {
            // Traverse Left 
            for j := colEnd; j >= colBegin; j-- {
                // res = append(res, matrix[rowEnd][j])
                matrix[rowEnd][j] = count 
                count++
            }
        }
        rowEnd--
        
        
        if colBegin <= colEnd {
            // fmt.Println(count)
            // Traver Up 
            for j := rowEnd;j >= rowBegin; j --{
                // res = append(res, matrix[j][colBegin])
                matrix[j][colBegin] = count
                count++
            }
        }
        colBegin++
    }
    

    return matrix
}