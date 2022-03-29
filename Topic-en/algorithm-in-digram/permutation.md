
# 字符串的排序顺序

- 手撸了一个方法，但是是错的。 原因是因为，没有考虑 去重 和排序的相关方案。 


```

func permutation(s string) []string {
    if len(s) == 0 {
        return nil 
    }
    // 可以选择的 [path], 当前路径 []
    // 回溯方法
    // 有没有递归的解法？ 
    // 用栈可能有
    return backtrack(s, []byte{}, []string{}, 0)
}
// 输入，输出 和 返回值的确认。还是有问题 
// 


func backtrack(s string, path []byte, res []string, curIdx int) []string {
    if len(s) == curIdx {
        // 加入

        // 这里是否需要去重？

        // 多叉树遍历方式。
        res = append(res, string(path))
        return res 
    }

    for i := curIdx ; i < len(s); i++ {   // ？
        path = append(path, s[i]) // add cur idx to path path 
        // recursive logic 
        // 有重复元素了。
        // 在哪一步会产生重复元素呢？
        res = append(res, backtrack(s, path, res, curIdx + 1 )...)
        path = path[:len(path)-1] // backtrack 
    }
    return res 
}

```



- 还是不怎么熟练


```

func permutation(s string) []string {
    if len(s) == 0 {
        return nil 
    }
    // 可以选择的 [path], 当前路径 []
    // 回溯方法
    // 有没有递归的解法？ 
    // 用栈可能有

    schar := []byte(s)
    res := []string{}
    return backtrack(schar, res, 0)
}
// 输入，输出 和 返回值的确认。还是有问题 
// 


func backtrack(schar []byte, res []string, curIdx int) []string {
    if len(schar) == curIdx {
        // 加入

        // 这里是否需要去重？

        // 多叉树遍历方式。
        // res = append(res, string(schar))
        return []string{string(schar)}
    }

    // 只需要每条分支去重即可
    set := map[byte]bool{}
    for i := curIdx ; i < len(schar); i++ {   
        if set[schar[i]] {
            continue // 重复元素了
        }

        set[schar[i]] = true 
        // path = append(path, schar[i]) // add cur idx to path path 
        // recursive logic 
        // 有重复元素了。
        // 在哪一步会产生重复元素呢？

        // 交换元素。 
        schar[i], schar[curIdx] = schar[curIdx], schar[i]
        res = append(res, backtrack(schar, res, curIdx + 1)...)

        backtrack(schar, res, curIdx + 1)
        // path = path[:len(path)-1] // backtrack 

        schar[i], schar[curIdx] = schar[curIdx], schar[i] // 恢复交换
    }
    return res 
}

```


```

func permutation(s string) (ans []string) {
    t := []byte(s)
    sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
    n := len(t)
    perm := make([]byte, 0, n)
    vis := make([]bool, n)
    var backtrack func(int)
    backtrack = func(i int) {
        if i == n {
            ans = append(ans, string(perm))
            return
        }
        for j, b := range vis { // 遍历 set 元素。 遍历的是 vis 数组....
            if b || j > 0 && !vis[j-1] && t[j-1] == t[j] { // 去重关键。 
                continue
            }
            vis[j] = true
            perm = append(perm, t[j])
            backtrack(i + 1)
            perm = perm[:len(perm)-1]
            vis[j] = false
        }
    }
    backtrack(0)
    return
}

```
