# 校验压入和弹出栈的相关队列


````
func validateStackSequences(pushed []int, popped []int) bool {
    stack := []int{}
    i := 0 

    for _, n := range pushed {
        stack = append(stack, n)
        for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
            stack = stack[:len(stack)-1]
            i += 1 
        }
    }
    
    if len(stack) > 0 {
        return false
    } else {
        return true 
    }
}
```

