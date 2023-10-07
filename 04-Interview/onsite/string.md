

```

package main

import (
    "fmt"
)
func main() {
    //a := 0
    //fmt.Scan(&a)
    //fmt.Printf("%d\n", a)
    fmt.Printf("Hello World!\n");
}
// 1 
// 输入
// [1, 3, 4, 5, 6] （固定）
// [3, 4, 5, 6, 7]
// [3, 5, 6]

// 输出
// [3, 4, 5, 6]


// 输出
// [3, 5, 6]

// 动态规划
// 双串。
// string1, string2 

// 状态 dp， 以 N 为结尾的 string 与 以 J 结尾的 string2 的公共子串。
// 思路，动态规划的状态转移表达式， 减少问题规模。
// string1[:N-1], string2[:J-1], 这两个子问题的状态转移过程。
// str1 :   
// dp 二维矩阵 —》

//      [1, 3, 4, 5, 6] str1 
// str2 
// 3. 0  0, 1, 1(*), 1, 1

// 5. 0, 0, 1, 1, 2(*), 2 

// 6, 0, 0, 1, 1, 2, 3(*)


// [3, 5, 6]

// 核心代码，状态转移代码
// if string1[i-1] == string2[j-1] {
//     dp[i][j] = dp[i-1][j-1] + 1 
//.    // 记录一下元素。
// } else {
//     dp[i][j] = maxe(dp[i-1][j], dp[i][j-1])
// }

// dp[i][j], 从最右下角开始。


func maxCommonString(s1 string, s2 string) string {
    m := len(s1)
    n := len(s2)

    // 初始化 dp 数组
    dp := make([][]int, m + 1)
    for i := 0; i < m + 1; i ++ {
        dp[i] = make([]int, n + 1)
    }
    
    // 填充 dp 
    // dp 最优解
    ret := []rune{}
    for i := 1; i < m + 1; i ++ {
        for j := 1; j < n + 1; j ++{
            if s1[i-1] == s1[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1 
                /// ret = append(ret, rune(s1[i-1]))
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    // dp 数组构建完之后，再推导一次。
    // dp 的右下角，往前进行推到，然后记录最大值。

    return string(ret)
}

func max(i, j int) int {
    if i > j {
        return i 
    } else {
        return j 
    }
}


```
