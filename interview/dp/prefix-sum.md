

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


