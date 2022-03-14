- 跳台阶

- jump floor 相关代码

```

func numWays(n int) int {
    // 1. 用斐波那契。 
    // 2. 用完全跳台阶的方法去解答。 [✅]
    if n == 0 || n == 1 {
        return 1 
    }

    // 用这种写法，而不用 【滚动数组】是因为这种写法更能够复用以后的代码。
    num := []int{1, 2}
    dp := make([]int, n + 1)
    dp[0] = 1 // base case dummy case 因为可以跳两次。 
    dp[1] = 1 // 台阶 1 
    
    for j := 2; j <= n; j ++ {
        for _, i := range num {
            dp[j] = (dp[j] + dp[j-i]) % 1000000007 // 取模
        }
    }
    return dp[n]
}

```
