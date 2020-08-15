package main 

func findUnsortedSubarray(nums []int) int {
    // selection sorting 的一个例子
    // 选择排序的一个性质来做的
    l := len(nums)
    r := 0 
    for i:=0;i < len(nums);i++{
        for j:=i+1;j<len(nums);j++{
            if nums[j] < nums[i] {
                r = max(r, j)
                l = min(l, i)
            }
        }
    }
    if r - l < 0 {
        return 0
    } else {
        return r - l + 1
    }
}

func max(a, b int) int {
    if a > b {
        return a 
    } else {
        return b 
    }
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a 
    }
}