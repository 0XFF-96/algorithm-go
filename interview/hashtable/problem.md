- 判断是否同构字符串

```

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false 
    }
    s2t := map[byte]byte{}
    t2s := map[byte]byte{}

    for i := range s {
        x, y := s[i], t[i]
        if (s2t[x] > 0 && s2t[x] != y ) || 
            (t2s[y] > 0 && t2s[y] != x ) {
            return false 
        }
        // 构建双射关系
        s2t[x] = y 
        t2s[y] = x 
    }
    return true 
}
```


- 

- 两个列表共同元素的最小索引和。 


```

func findRestaurant(list1 []string, list2 []string) []string {
    // 两个列表的最小索引总和

    minIdx := 10000000
    ret := make([]string, 0)
    commonSet := map[string]int{}
    for idx, i := range list1 {
        commonSet[i] = idx
    }
    
    for idx, i := range list2 {
        if commonIdx, ok := commonSet[i]; ok {
            if commonIdx + idx == minIdx {
                ret = append(ret, i)
            }
            if commonIdx + idx < minIdx {
                minIdx = commonIdx + idx 
                ret = []string{i} // 抛弃，然后重建
            }
        }
    }
    return ret
}

```
