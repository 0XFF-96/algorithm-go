
# LegnthOfLogestSubstring 

一共有三种方法做 


- 双指针 + 哈希表

```

func lengthOfLongestSubstring(s string) int {
    dic := map[byte]int{}
    res := 0 
    i := -1 // 左边界 
    for j := 0; j < len(s); j++{
        if _, ok := dic[s[j]]; ok {
            i = max(dic[s[j]], i) // 更新左指针
        }
        dic[s[j]] = j // 最新记录
        res = max(res, j - i ) // 更新结果， 为什么不用 - 1 
    }
    return res 
}

```

- 动态规划 + 线性查找 

```

func lengthOfLongestSubstring(s string) int {
    res := 0 
    tmp := 0 
    i := 0 

    for j := 0; j < len(s); j ++ {
        i = j - 1 // why ? 初始化
        for i >= 0 && s[i] != s[j] {
            i -= 1 // 线性查找 i 
        }

        if tmp < j - i {
            tmp = tmp + 1 // 没搞懂 ？dp[j - 1] -> dp[j]
        } else {
            tmp = j - i 
        }

        res = max(res, tmp) // max dp[j], d[j-1]
    }
    return res 
}

```
