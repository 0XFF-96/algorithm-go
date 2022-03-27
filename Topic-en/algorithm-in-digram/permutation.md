
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
