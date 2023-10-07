

- 统计不同城市的路径数量

```

func countRoutes(locations []int, start int, finish int, fuel int) int {
    // locations = [2, 3, 6, 8, 4], start = 1, finish = 3, fuel = 5
    // 城市          0, 1, 2,3 , 4


    // 输入：locations = [4,3,1], start = 1, finish = 0, fuel = 6
    // 输出：5
    
    // 二维 dp 表，能否用来搜索？
            2, 3, 6, 8, 4 
    // 0,   0  1  4, 6, 2
    // 1,   1  0 
    // 2,  
    // 3,  
    // 4,    

    // 上面的表，只表达了，两两城市之间的相互关系。 

    // 如何表达， 1 - 3 城市的方法？

    // 1. 减少问题规模 ,  [2, 3] -> [2, 3, 6] , 从 2 个城市，到 3 个城市，会有什么区别吗？
    // 2. 子问题之间的联系。
    // 对角线应该都是 0 , 因为自己到自己城市，不会消耗 fuel 。 


    // 思考方法
    // 特定「起点」，明确且有限的「移动方向」（转移状态），求解所有状态中的最优值。
    // 没有特定的移动方法， 感觉就是一个 DFS 的方法。 
    // 把所有城市，视作一个【互相连通】 的图。 用 DFS + 回溯的方法

    // 解题方法1
    // 1. 递归方法， 复杂度很高。但是代码相对可读。 
    // 2. 迭代方法， 记忆化搜索，可以利用空间换时间的方法，提前计算。
}

```

- 最差的 golang 实现（） 照着模板改的。 
- 在 need < fuel, need <= fuel, fuel > need  的条件判断改写中，出了错。耽误了一些时间进行 debug 处理。 

```
var mod = 1000000007
var cache [][]int  // 为什么需要用这个 cache ?
// 二维数组每一个代表什么意思？

func countRoutes(locations []int, start int, finish int, fuel int) int {
    if len(locations) == 0 {
        return 0 
    }

    n := len(locations)
    cache = make([][]int, n)
    
    for i := 0; i < len(cache); i ++ {
        cache[i] = make([]int, fuel + 1) // fuel + 1 避免不必要判断？！
    }

     // 初始化 cache 矩阵
    for i := 0; i < n; i++ {
        for j := 0; j < fuel + 1; j++{
                cache[i][j] = -1 
        }
    }

    // fmt.Println(cache)

    // return: 在位置 u 出发， 油量为 fuel 的前提下
    // 到达 end 的 [ 路径数量 ]
    return dfs(locations, start, finish, fuel)
}

func dfs(ls []int, u, end, fuel int) int {
    // <<<<<<<<<<<<<<<<<< 递归基开始 >>>>>>>>>>>>>>>>>>>>>>>
    // 递归结束条件？
    // 有数值，说明已经走过
    // 其实 cache 用零初始化也是可以的对吧。
    if cache[u][fuel] != -1 {
        return cache[u][fuel]
    }

    // 先判断是否够 fuel
    // 重要性质，如果不能一步到终点到话，后续也不能到终点了。
    need := abs(ls[u] - ls[end])
    // 1. 不能到达的情况
    if need > fuel { // 这里为什么不能等于？ 因为后面可以判断
        cache[u][fuel] = 0 
        return 0 
    }

    // fmt.Println("dfs:", u, end, fuel)

    n := len(ls)
    // 开始 dfs 
    // 类似于回溯的算法，但是不需要 backtrack 
    // 因为有 cache 记录了是否已走过。
    sum := 0 
    // 如果已经是可以行的路径，就先算 1 
    if u == end {
        sum = 1 
    }

    // <<<<<<<<<<<<<<<<<< 递归基结束 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>

    // 1. 从 locations 数组开始枚举 到 u 这个点的距离。
    // 2. 逐步消耗 fuel 
    for i := 0; i < n; i++ {
        // i = u 时无须枚举，否则会死循环
        if i != u {
            need := abs(ls[i] - ls[u])
            if  need <= fuel { // need > fuel, 和上面一致的判断。 // need < fuel , fuel >= need 
                // 能否用多叉树结构，表示这个递归深入的过程？
                sum += dfs(ls, i, end, fuel - need) // 递归深入
                sum %= mod 
            }
        }
    }
    cache[u][fuel] = sum
    return sum 
}


func abs(i int) int {
    if i > 0 {
        return i 
    } else {
        return -i
    }
}
```
