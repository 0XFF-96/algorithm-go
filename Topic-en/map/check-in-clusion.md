
# check in clusion

```

func checkInclusion(s1 string, s2 string) bool {
    // Obviously, brute force will result in TLE. Think of something else.  
    // How will you check whether one string is a permutation of another string?
    // One way is to sort the string and then compare. But, Is there a better way?
    // If one string is a permutation of another string then they must one common metric. What is that?
    // Both strings must have same character frequencies, if one is permutation of another. Which data structure should be used to store frequencies?
    // What about hash table? An array of size 26?


    n, m := len(s1), len(s2)
    if n > m {
        return false 
    }

    var cnt1, cnt2 [26]int

    for i, ch := range s1 {
        cnt1[ch - 'a']++
        cnt2[s2[i] - 'a']++
    }

    if cnt1 == cnt2 {
        return true 
    }

    for i := n; i < m; i++ {
        cnt2[s2[i] - 'a']++ // 滑动窗口增加
        cnt2[s2[i-n] - 'a']-- // 滑动窗口减少
        if cnt1 == cnt2 {
            return true 
        }
    }

    return false 
}
```


