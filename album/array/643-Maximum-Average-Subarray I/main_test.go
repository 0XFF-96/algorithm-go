package main 

// pacakge: https://leetcode.com/problems/maximum-average-subarray-i/solution/
func findMaxAverage(nums []int, k int) float64 {
    // 滑动窗口的实例
    // 
    // 
    // two points 
    // 左指针每移动一步，减去一个数
    // 右指针每移动一步，加上一个数
    sum := 0 
    for i:=0;i<k;i++{
        sum += nums[i]
    }
    res := sum
    
    for i:=k;i<len(nums);i++{
        sum += nums[i] - nums[i-k]
        if sum > res {
            res = sum
        }
    }
    return float64(res)/float64(k)
}

// 方法二： Approach #1 Cumulative Sum [Accepted]
// https://leetcode.com/problems/maximum-average-subarray-i/solution/