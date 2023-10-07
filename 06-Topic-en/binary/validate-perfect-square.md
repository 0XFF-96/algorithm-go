# Validate perfect square 

- 不用 sqart() 库函数， 但是又需要算平方的题目。 一般用二分查找就很值得。 

```
func isPerfectSquare(num int) bool {
    return isPerfect(1, num, num)   
}

func isPerfect(low, high, num int) bool {
    if low > high {
        return false
    }
	
    mid := (low + high) / 2
    if mid * mid == num {
        return true
    }
    if mid * mid < num {
        return isPerfect(mid + 1, high, num)
    }
    return isPerfect(low, mid - 1, num)
}

```


```

func isPerfectSquare(num int) bool {
    if num == 1 {
        return true
    }
    
    start, end := 0, num
    var mid int
    for  start <= end {
        mid = (start + end + 1) / 2
        if p := mid * mid; p == num {
            return true
        } else if p > num {
            end = mid - 1
        } else {
            start = mid + 1
        }
    }
    return false
}

```
