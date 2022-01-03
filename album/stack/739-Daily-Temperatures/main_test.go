func dailyTemperatures(T []int) []int {
    // for-for 
    // next greater element 1. or 2 
    res := make([]int, len(T))
    stack := []int{}
    
    for i:= 0; i < len(T); i ++ {
        // 单调栈
        for len(stack) != 0 && T[stack[len(stack)-1]] < T[i] {
            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res[idx] = i - idx
        }
        stack = append(stack , i)
    }
    return res
}