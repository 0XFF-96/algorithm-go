
- 7 进制转换

```

func convertToBase7(num int) string {
    if num == 0 {
        return "0"
    }

    isMinus := false 
    var ans string
    if num < 0 {
        isMinus = true 
        num = -num // 倒转正负
    }

    for num != 0 {
        // 本来应该是倒序取的，用头插法就可以免了。
        ans = strconv.Itoa( num % 7) + ans 
        num /= 7 
    }

    if isMinus {
        ans = "-" + ans 
    }
    return ans
}

```

- 海明距离，位为 1 的个数。

```

func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num &= (num-1)
		cnt++
	}
	return cnt
}

```

- 颠倒二进制

```
func reverseBits(n uint32) (rev uint32) {
    for i := 0; i < 32 && n > 0; i++ {
        rev |= n & 1 << (31 - i)
        n >>= 1
    }
    return
}

```

- 不用加法的两数之和。 

```

func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	return getSum(a^b, ((a&b) << 1))
}

```
