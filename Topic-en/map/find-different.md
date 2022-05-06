
# find different 

- 389. Find the Difference


```

func findTheDifference(s, t string) byte {
    cnt := [26]int{}
    for _, ch := range s {
        cnt[ch-'a']++
    }
    for i := 0; ; i++ {
        ch := t[i]
        cnt[ch-'a']--
        if cnt[ch-'a'] < 0 {
            return ch
        }
    }
}

```


```

func findTheDifference(s, t string) (diff byte) {
    for i := range s {
        diff ^= s[i] ^ t[i]
    }
    return diff ^ t[len(t)-1]
}

```
