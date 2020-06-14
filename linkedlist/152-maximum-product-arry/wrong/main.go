func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }
    if len(nums) <= 1 {
        return nums[0]
    }
    // 动态规划
    // 分解为子表达式的值
    // subarray 和 sub sequence 之间的区别。
    // 
    
    var curMax int 
    var ret int 
    // 0 1 0 1 的序列
    // The result cannot be 2, because [-2,-1]
    // we have recursive solution about this question
    // f(n) = max(f(n-1)*current or f(n-2)*current)  not 
    // 
    // f(n) = max(current, f(n-1)*currrent)    这才是子问题的解法
    
    
    // 有一种 case 过不了
    // 看看是为什么。
    // [-2,3,-4]
    // [2,3,-2,4]
    
    // 记乘积，最小的数目
    min := 1
    curMax = 1
    for _, n := range nums {
        curMax = max(n, n*curMax)
        if curMax > ret {
            ret = curMax
        }
    }
    return ret
}

func max(i, j int) int {
    if i > j {
        return i
    } else {
        return j
    }
}
