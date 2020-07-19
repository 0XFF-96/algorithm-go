


func validateStackSequences(pushed []int, popped []int) bool {
    if len(pushed) != len(popped) {
        return false
    }
    // 1. pushed is a permutaion of popped 
    // 2. have distinct values 
    
    // 模拟入栈操作
    // 关键点在于，如何移动两个头的指针。 
    // 什么情况是失败的？能够提前返回吗？    
    // 怎么遍历这两个栈？

    j := 0 
	stack := []int{}
	
	// 为什么是从 pushed 栈开始，
	// 而不是从 poped 栈开始解题目？
    
    for _, s := range pushed {
        stack = append(stack, s)
        
        for len(stack) != 0 && j < len(popped) && stack[len(stack)-1] == popped[j] {
            // popped
            stack = stack[:len(stack)-1]
            j += 1
        }
    }
    return j == len(popped)
}

// 还有没有其他解法？
