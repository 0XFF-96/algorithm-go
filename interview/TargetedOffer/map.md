

- 最长不重复子串， （不是子序列） 

```
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {return 0 }
    res := 0 
    set := map[byte]int{}
    start := 0 
    end := 0 

    // 如果有中文字符，就不能这样做了。 
    // rune 和 utf-8 的原因
    for end < len(s) {
        alpha := s[end]
        if _, ok := set[alpha]; ok {
            start = max(set[alpha], start) // 快速移动
        }

        res = max(res, end - start + 1)
        set[s[end]] = end + 1 
        end++
    }
    return res 
}

func max(i, j int) int {
    if i > j {
        return i
    } else {
        return j 
    }
}
```
