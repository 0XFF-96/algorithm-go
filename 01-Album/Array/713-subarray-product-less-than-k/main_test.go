package main 
// 基本 AC 
// 但是遇到数据量大的情况，不行。
func numSubarrayProductLessThanK(nums []int, k int) int {
    // 排序->剪枝
    // 2， 5， 6， 10 
    // 什么是 sub array ？不能够排序，因为是 sub array ，而不是 sub sequence
    // 如何去重？
    // 类型：利用 dfs 相关技巧解题目
    // 如何进行 early stop
    
    // 1 个元素， 2 个元素， 3 个元素
    // 两个 for 循环？
    
    // edge case 全部都是 1 的情况，
    // 这种输入不过。
    // 输入太大了
    count := 0 
    for i:=0;i < len(nums); i ++ {
        pSum := nums[i] 
        for j :=i+1;j < len(nums); j ++ {
            if pSum < k {
                count++
            } else {
                break
            }
            pSum = pSum * nums[j]
        }
        if pSum < k {
            count++
        }
    }
    return count
}

// accept 
// Runtime: 128 ms, faster than 78.85% of Go online submissions for Subarray Product Less Than K.
// Memory Usage: 7.6 MB, less than 44.23% of Go online submissions for Subarray Product Less Than K.
// 这是利用滑动窗口的题目
func numSubarrayProductLessThanKV2(nums []int, k int) int {
    if k <= 1 {
        return 0 
    }
    ans := 0 
    left := 0 
    prod := 1
    
    for right := 0; right < len(nums); right++{
        prod *= nums[right]
        
        for prod >= k {
            prod = prod / nums[left]
            left += 1
        }
        ans += right - left + 1
    }
    return ans
}
