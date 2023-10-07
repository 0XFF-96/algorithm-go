
# 开根号


```

func mySqrt(x int) int {
    r := x 
    for r * r > x {
        r = (r + x/r) /2 
        // 为什么能够用，类似二分的方法呢？
        // 开根号的算法。
        // r/2 + x/2r
        // (8 + 8/8) /2 = 4+1
        //  (5 + 8/5) /2 
    } 
    return r 
}

```

- 二分法。 

```

```

# 367 有效数的完全平方

```
func isPerfectSquare(num int) bool {
    left, right := 0, num
    for left <= right {
        mid := (left + right) / 2
        square := mid * mid
        if square < num {
            left = mid + 1
        } else if square > num {
            right = mid - 1
        } else {
            return true
        }
    }
    return false
}

```
