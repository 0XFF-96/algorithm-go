



// 接雨水问题


// 单调栈解法
func trap(height []int) (ans int) {
    stack := []int{}
    for i, h := range height {
        for len(stack) > 0 && h > height[stack[len(stack)-1]] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                break
            }
            left := stack[len(stack)-1]
            curWidth := i - left - 1
            curHeight := min(height[left], h) - height[top]
            ans += curWidth * curHeight
        }
        stack = append(stack, i)
    }
    return
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}