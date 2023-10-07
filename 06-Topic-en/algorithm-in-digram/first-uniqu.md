
# First unique char


```

func firstUniqChar(s string) byte {
    var arr = make([]int, 26)
    length := len(s)
    for i :=0 ; i < length; i++{
        arr[s[i]-'a']++
    }
    for i := 0; i < length; i++{
        if arr[s[i]-'a'] == 1 {
            return s[i]
        }
    }
    return ' '
}

```
