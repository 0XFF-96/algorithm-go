- 字母异位词。

```

func groupAnagrams(strs []string) [][]string {
    // key = 按照字母排序后的设计
    // s string -> []rune 的转换。需要排序。 

    mp := map[string][]string{}
    for _, str := range strs {
        s := []byte(str)
        sort.Slice(s, func(i, j int) bool {
            return s[i] < s[j]
        })
        sortedStr := string(s)
        mp[sortedStr] = append(mp[sortedStr], str)
    }

    ans := make([][]string, 0, len(mp))
    for _, v := range mp {
        ans = append(ans, v)
    }
    return ans 
}

```
