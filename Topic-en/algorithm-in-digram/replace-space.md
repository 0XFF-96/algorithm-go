
- 替换空格

```

func replaceSpace(s string) string {
    b := []byte(s)
    result := make([]byte, 0)

    for i := 0; i < len(b); i ++ {
        if b[i] == ' ' {
            result = append(result, []byte("%20")...)
        } else {
            result = append(result, b[i])
        }
    }
    return string(result)
}

```


```

// 原地修改版本
func replaceSpace(s string) string {
    b := []byte(s)
    oldlength := len(b)

    spaceCount := 0 
    // 1. 计算空格数量
    for _, v := range b {
        if v == ' ' {
            spaceCount++
        }
    }

    // 2. 扩展原有切片 
    resizeCount := spaceCount * 2 
    tmp := make([]byte, resizeCount)
    b = append(b, tmp...) // 应该是另一种更高效的方法？ 例如可以直接 copy 

    // 双指针
    i := oldlength - 1 // 之前的长度
    j := len(b) - 1    // 新的长度

    for i >= 0 {
        if b[i] != ' ' {
            b[j] = b[i]
            j--
        } else {
            b[j] = '0'
            b[j-1] = '2'
            b[j-2] = '%'
            
            j = j - 3 
        }
        i--
    }
    return string(b)
}

```
