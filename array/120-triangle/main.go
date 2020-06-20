// func minimumTotal(triangle [][]int) int {
//     // from top to bottom 
//     // from bottom to top 
//     // limitation: move to adjacent numbers on the row below 
//     // adjacent numbers is xxxx 
//     //
    
//     // 算法：从底向上遍历，
//     // 贪心法：取 adjacent numbers 最小的那个
//     // 最后 for 循环一遍，找出最小元素
//     // O(n) 和 O(n) 的算法
    
    
//     // 上面想法是错的
//     // 原因在于？
//     // 有可能路径不太对
//     // 需要用 DP 的思想来做
// }


func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	var dp [][]int = make([][]int,rows)
	// 初始化  triangle
    for i := 0;i < rows;i++{
		dp[i] = make([]int,i + 1)
	}
    
    // 赋值初始化
	for i := 0;i < rows;i++{
		dp[rows-1][i] = triangle[rows-1][i]
	}
	for i := rows - 2;i >= 0;i--{
		for j := 0;j <= i;j++{
			dp[i][j] = triangle[i][j] + min_int(dp[i + 1][j],dp[i + 1][j + 1])
		}
	}
	return dp[0][0]
}

func min_int(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}