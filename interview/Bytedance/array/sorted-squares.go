package 



// 双指针的写法

// 同样地，我们可以使用两个指针分别指向位置 00 和 n-1n−1，
// 每次比较两个指针对应的数，选择较大的那个逆序放入答案并移动指针。这种方法无需处理某一指针移动至边界的情况，读者可以仔细思考其精髓所在


func sortedSquares(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)   // 这个修改。其实还可以进一步压缩空间的，直接在给出来的 nums 数组上，进行修改。
    i, j := 0, n-1
    for pos := n - 1; pos >= 0; pos-- {
        if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
            ans[pos] = v
            i++
        } else {
            ans[pos] = w
            j--
        }
    }
    return ans
}

