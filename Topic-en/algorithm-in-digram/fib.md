
# fib 斐波那契数量

- 解法一。 

```

func fib(n int) int {
    if n == 0 {
        return 0 
    } 

    if n == 1 {
        return 1 
    }

    return fib(n - 1) + fib(n - 2)
}

```

- 解法二。 

``` 

func fib(n int) int { 
    dp := make([]int, n + 1) 
    return fibMemorized(n, dp)
}

func fibMemorized(n int, dp []int) int {
    if n == 0 {
        return 0 
    }

    if n == 1 {
        return 1 
    }

    dp[n] = fibMemorized(n - 1, dp) + fibMemorized(n - 2, dp)
    return dp[n]
}

```

- 解法三。进一步优化，动态规范的方法。 


```
func fib(n int ) int {
    if n == 0 {
        return 0 
    }

    if n == 1 {
        return 1 
    }

    dp := make([]int, n + 1) // 为什么需要加1
    dp[1] = 1 

    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }

    return dp[n]
}
```

- 解法三。 上面错误是应为没有取模。 % 
- 而且需要在最后才取模，而不是一开始。 

```

func fib(n int ) int {
    if n == 0 {
        return 0 
    }

    if n == 1 {
        return 1 
    }
    dp0 := 0
    dp1 := 1

    for i := 2; i < n + 1; i ++ {
        tmp := ( dp0 + dp1 ) % 1000000007
        dp0 = dp1 
        dp1 = tmp 
    }
    return dp1 
}

```
