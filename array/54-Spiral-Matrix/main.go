package main 

// input: [[1,2,3],[4,5,6],[7,8,9]]
// outpu: [1,2,3,6,9,8,7,4,5] 
// 只要搞明白，这代码和上面输入输出对应的关系就好了

func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return nil 
    }
    r1 := 0 
    r2 := len(matrix) -1 
    c1 := 0 
    c2 := len(matrix[0]) -1
    
    var ans []int
    for r1 <= r2 && c1 <= c2 {
        // 固定 r1
        for c :=c1; c <= c2;c++{
            ans = append(ans, matrix[r1][c])
        }
        // 固定 c2
        for r := r1+1; r <= r2; r++ {
            ans = append(ans, matrix[r][c2])
        }        
        
        // 什么情况不满足这条件？
        if (r1 < r2 && c1 < c2) {
            
            // 固定 r2
            for c:= c2-1; c > c1; c-- {
                ans = append(ans, matrix[r2][c])
            }
            
            // 固定 c1 
            for r := r2; r > r1; r-- {
                ans = append(ans, matrix[r][c1])
            }
        }
        r1++;
        r2--;
        c1++;
        c2--;
    }
    return ans
}