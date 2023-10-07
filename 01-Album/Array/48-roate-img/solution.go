// 思路
// 交换对于下标的元素、逆序。
// 
func rotate(matrix [][]int)  {
    // n*n 
    row := len(matrix)  
    column := len(matrix[0]) 
    
    for r := 0; r < row; r++{
        for c:= 0; c < column; c ++{
            // if c == r {
            //     continue 
            // }
            // 被换了两次，然后又换回来了...
            
            // 应该要联合 r
            if r >= c {
                continue 
            }
            matrix[r][c],matrix[c][r] = matrix[c][r],matrix[r][c]
        }
        
        // 这里不应该加 = 号..
        // 偶数的时候，不需要换.
        for c:=0; c < column/2; c++{
            matrix[r][c], matrix[r][column-c-1] = matrix[r][column-c-1], matrix[r][c]
        }
        // 数组逆序
    }
    // 一步到位的计算方法。
    // 
    // 0， 1 -> 
    // 1， 0 
}