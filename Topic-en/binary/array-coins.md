
# Arrange coins 


```
 func arrangeCoins(n int) int {
    // 斐波那契？
    // 1 = 1 
    // 2 = 2 
    // 3 = 3 
    cnt := 0 
    count := 0
    for cnt < n {
        cnt += count + 1 
        if cnt > n {
            break
        }
        count++
    }
    return count 
}
```

- 利用二分法。
- 利用数学公式。 

```
func arrangeCoins(n int) int {
    // 斐波那契？
    // 1 = 1 
    // 2 = 2 
    // 3 = 3 
    // Why so many down votes on this question?
    // It is a really good binary search problem.
    
    left, right := 0, n 
    for left <= right {
        k := (right + left) / 2 
        curr := (k * (k + 1) )/ 2 
        
        if curr == n {
            return k 
        }
        
        if n < curr {
            right = k - 1 
        } else {
            left = k + 1 
        }
    }

    return right 
}
```
