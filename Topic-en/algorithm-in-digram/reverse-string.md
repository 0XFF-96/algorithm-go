

- 把字符串左移 N 位 。 


```

func reverseLeftWords(s string, n int) string {
    b := []byte(s) 
    // 1. 反转前 n 个字符
    // 2. 反转 n 到 end 个字符
    // 3. 反转整个字符
    reverse(b, 0, n-1)
    reverse(b, n, len(b) - 1)
    reverse(b, 0, len(b) - 1)
    return string(b)
}

func reverse(b []byte, left, right int) {
    for left < right {
        b[left], b[right] = b[right], b[left]
        left++
        right--
    }
}

```
