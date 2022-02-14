

- 前缀和

```



func Constructor(nums []int) NumArray {
    if len(nums) == 0 {
        return NumArray{}
    }
    sum := make([]int, len(nums) + 1)

    // 重点！
    // 我们插入了一个虚拟 0 作为 sum 数组中的第一个元素。
    // 这个技巧可以避免在 sumrange 函数中进行额外的条件检查
    sum[0] = nums[0]
    // for i:= 1; i <= len(nums) ; i++{ 
    //     sum[i ] = sum[i-1] + nums[i-1]
    // }

    for i := 0; i < len(nums); i++ {
        sum[i + 1] = sum[i] + nums[i]
    }

    // 有两种初始化前缀和的策略。



    // sum(i,j) = sum(0,j+1) - sum(0,i)
    fmt.Printf("%+v", sum)
    return NumArray{prefixSum: sum}
}


func (this *NumArray) SumRange(left int, right int) int {
    // 检查 right 和 left 是否正确
    // [) , 这是一个左开右闭的区间。
    // 所以 right - 1
    // left <= right

    // 区间查询方式。
    // sum(i, j) = sums(0, j + 1) - sum(0, i)
    return this.prefixSum[right + 1] - this.prefixSum[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

 ```

- 二维前缀和的构建

```

type NumMatrix struct {
    dp [][]int 
}

func Constructor(matrix [][]int) NumMatrix {
    if len(matrix) == 0 {
        return NumMatrix{}
    }
    dp := make([][]int, len(matrix) + 1) // 注意需要加一，前缀和特征。
    for i := 0; i < len(matrix) + 1; i ++ {
        dp[i] = make([]int, len(matrix[0]) + 1)
    }

    // 构造二维前缀和
    for r := 1; r < len(matrix) + 1 ; r++ {
        for c := 1; c < len(matrix[0]) + 1; c++{
            dp[r][c] = dp[r-1][c] + dp[r][c-1] + matrix[r-1][c-1] - dp[r-1][c-1]
        }
    }

    // 另一种，构造二维前缀和的方式。r+1, c+1 开始的

    return NumMatrix{dp: dp}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    // 这个公式可能记不住的，怎么办？
    // 怎么记住？

    // 记住题目的提示关系。
    // 0 <= row1 <= row2 < m
    /// 0 <= col1 <= col2 < n
    return this.dp[row2 + 1][col2 + 1] - this.dp[row1][col2+1] - this.dp[row2 +1][col1] + this.dp[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

 ```
