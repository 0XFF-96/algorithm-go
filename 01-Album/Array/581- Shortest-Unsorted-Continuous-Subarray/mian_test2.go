package main 

// 利用排序优化这个算法。
// 
func findUnsortedSubarrayV2(nums []int) int {
    // 暴力解法
    snums := append([]int{}, nums...)
    sort.Ints(snums)
    
    start := len(snums)
    end := 0 
    
    for i :=0; i < len(snums); i ++ {
        if snums[i] != nums[i] {
            start = min(start, i)
            end = max(end, i)
        }
    }
    if end - start >= 0 {
        return end - start + 1
    } else {
        return 0 
    }
}



// Approach 4: Using Stack ! s 
func findUnsortedSubarrayV3(nums []int) int {
    // 暴力解法
    stack := []int{}
    l := len(nums)
    r := 0 
    
    for i :=0; i < len(nums); i++ {
        for len(stack) != 0 && nums[stack[len(stack)-1]] > nums[i] {
            l = min(l, stack[len(stack)-1])
            stack = stack[:len(stack)-1]
        }
        // 为什么 push 的是 Index 
        stack = append(stack, i)
    }
    // clear stack 
    stack = []int{}
    for i := len(nums) -1; i >= 0; i-- {
        for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
            r = max(r, stack[len(stack)-1])
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    
    if r - l > 0 {
        return r - l + 1
    } else {
        return 0 
    }
}