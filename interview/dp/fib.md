
- 数组中最长斐波那契数列数量的长度
- 可以用两种解题方法， 动态规划， 模拟斐波那契数列
```
func lenLongestFibSubseq(arr []int) int {
    // 利用 hash set 
    // 判断 f(n) = f(n-1) + f(n-2) 是否存在
    set := make(map[int]struct{})
    for _, i := range arr {
        set[i] = struct{}{}
    }

    ans := 0

    // 自底向上
    // 模拟斐波那契数列，判断是否存在 set 中
    // 没有用动态规划的思想解决题目。

    for i:=0; i < len(arr); i++ {
        for j:= i+1; j < len(arr); j++{
            x := arr[j]
            y := arr[i] + arr[j]
            length := 2 

            _, hasFabo := set[y]
            for hasFabo {
                // x, y -> y, x+y

                // 滚动数组，模拟斐波那契数列
                tmp := y 
                y += x 
                x = tmp 
                
                length++
                ans = max(ans, length)
                _, hasFabo = set[y]
            }
        }
    }

    // 根据斐波那契数列的定义。
    if ans >= 3 {
        return ans 
    } else {
        return 0 
    }
}
```
