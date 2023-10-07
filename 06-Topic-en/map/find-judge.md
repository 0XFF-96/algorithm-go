# Find judge 


```

// func findJudge(n int, trust [][]int) int {
//     // check rules 
//     // 图？ 
//     // 哈希表？
//     // judge = 没有出度，但是入度 = n 
//     // 

//     set := map[int]int{}

//     // row := len(trust)
//     // col := len(trust[0])
//     // for r :=0; r < row; r ++ {
//     //     for c :=0; c < col; c++ {
//     //     }
//     // }

//     for i := 0; i < len(trust); i ++ {
//         item := trust[i]
//         set[item[1]]++
//     }

//     // 1. 利用哈希表，建立 信任关系图， 出度、入度。

//     // 2. 检查 哈希表， 出度、入度。
//     // 3. 找出， 入度为 n-1, 出度为 0 的元素。 
//     // 4. 如果没有，则返回 0 。

//     // for i :=1 ; i <=n; i ++ {} 
//     // 计算各节点的入度和出度
// }


func findJudge(n int, trust [][]int) int {
    inDegrees := make([]int, n+1)
    outDegrees := make([]int, n+1)
    for _, t := range trust {
        inDegrees[t[1]]++
        outDegrees[t[0]]++
    }
    for i := 1; i <= n; i++ {
        if inDegrees[i] == n-1 && outDegrees[i] == 0 {
            return i
        }
    }
    return -1
}
```


