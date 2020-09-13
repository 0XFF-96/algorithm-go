package main 

func findLengthOfLCIS(nums []int) int {
    ans := 0 
    anchor := 0 
    
    for i:=0;i < len(nums);i++{
        if i > 0 && nums[i-1] >= nums[i] {
            anchor = i
        }
        ans = max(ans, i - anchor + 1)
    }
    return ans 
}

func max(a, b int) int {
    if a > b {
        return a 
    } else {
        return b 
    }
}