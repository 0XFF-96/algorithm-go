package main

func uniquePaths(m int, n int) int {
    // 类似动态规划的解法
    // 二维数组
    // 将行和列初始化为 1
    // 到达当前格子，有哪几种途径？
    // 
    // 考点1: 二维数组的遍历
    // 考点2: 分解问题的方法（比动态规划简单）
    // 构造二维数组，不熟练..
    // 
    // array := make([][]int, m)
    var array [][]int
    for i:= 0; i < n; i++{
        array = append(array, make([]int, m))
    }
    
    // 行
    for i := 0; i < m ;i ++{
        array[0][i] = 1
    }
    
    // 列
    for i :=0; i < n; i++{
        array[i][0] = 1
    }
    
    // 开始的方式不对？
    // 应该要初始化才对
    for i := 1; i < n; i++{
        for j:=1; j < m; j++{
            cur := 0 
            if j - 1 >= 0 {
                cur += array[i][j-1]
            }
            if i - 1 >= 0 {
                cur += array[i-1][j]
            }
            // 这里要不要加 1 操作？
            array[i][j] = cur  
        }
    }
    return array[n-1][m-1]
}