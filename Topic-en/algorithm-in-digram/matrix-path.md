
- 矩阵中的路径问题


```

func exist(board [][]byte, word string) bool {
    rows := len(board)  
    cols := len(board[0])  // rows and cols 是需要减的

    for i := 0; i < rows; i ++ {
        for j := 0; j < cols; j ++ {
            // 👆， 每个位置都启动一次 dfs 搜索。 
            // 1， 怎么优化这个？
            if (dfs(board, word, i, j, 0)) { 
                return true 
            }
        }
    }
    return false 
}


func dfs(board [][]byte, word string, i int, j int,  k int) bool {
    if i >= len(board)  || i < 0 {
        return false 
    }

    if j >= len(board[0])  || j < 0 {
        return false 
    }

    if board[i][j] != word[k] {
        return false 
    }
    
    if k == len(word) - 1  {
        return true 
    }


    // 👆 base case 和 递归结束条件

    board[i][j] = '*'

    res := dfs(board, word, i + 1, j, k + 1) ||
          dfs(board, word, i - 1, j, k + 1) ||
          dfs(board, word, i, j + 1, k + 1) ||
          dfs(board, word, i, j - 1, k + 1)

    // backtrack 
    board[i][j] = word[k]

    return res       
}


```
