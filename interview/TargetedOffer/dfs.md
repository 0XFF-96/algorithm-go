


- 矩阵中的路径 （dfs) 
- 利用 tmp 变量, 节省 visied map 的空间。

```
func exist(board [][]byte, word string) bool {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j ++ {
            if dfs(board, word, i, j, 0) {
                return true 
            }
        }
    }
    return false 
}

func dfs(board [][]byte, word string, x, y, curPosOfWord int) bool {
    if x < 0 || x >= len(board) {
        return false 
    }
    if y < 0 || y >= len(board[0]) {
        return false 
    }
    // 判断是否已经访问过
    if board[x][y] != word[curPosOfWord] {
        return false 
    }
    // 递归终止条件

    if curPosOfWord == len(word) -1 {
        return true 
    }

    // start backtrack 
    // 先标记
    tmp := board[x][y]
    board[x][y] = '/'
    
    result := dfs(board, word, x + 1, y, curPosOfWord +1) || 
              dfs(board, word, x - 1, y, curPosOfWord +1) ||
              dfs(board, word, x, y - 1, curPosOfWord +1) ||
              dfs(board, word, x, y + 1, curPosOfWord +1)

    // resume the backtrack value 
    board[x][y] = tmp 
    return result 
}
```

- moving count 
- 机器人的运动范围

```
func movingCount(m int, n int, k int) int {
    visited := make(map[int]bool)
    return dfs(0, 0, m, n, k, visited)
}

func dfs (i, j, m, n, k int, visited map[int]bool) int {
    if i < 0 || i >= m {
        return 0 
    }

    if j < 0 || j >= n {
        return 0 
    }

    if (i/10 + i %10 + j/10 + j%10) > k {
        return 0 
    }

    if visited[i + j * 1000] {
        return 0 
    }
    // hack 一下
    visited[i + j * 1000] = true 

    // 只需要向右下角移动，就可以了
    return dfs(i + 1, j, m, n, k, visited) + 
           dfs(i , j + 1, m, n, k, visited) + 1 
}
```
