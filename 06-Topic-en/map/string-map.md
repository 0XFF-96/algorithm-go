# String map 

```

func firstUniqChar(s string) int {
    cnt := [26]int{}
    for _, ch := range s {
        cnt[ch-'a']++
    }
    for i, ch := range s {
        if cnt[ch-'a'] == 1 {
            return i
        }
    }
    return -1
}
```
