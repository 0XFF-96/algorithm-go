package isPOPOrder

func IsPopOrder(pushStack []int, popStack []int) bool {
	// 在遍历完一遍入栈序列后，如果发现辅助栈不为空，则说明不能完全匹配弹出序列
	if len(pushStack) == 0 || len(popStack) == 0 || len(pushStack) != len(popStack){
		return false
	}
	var stack []int
	var popIndex int
	for i:=0; i < len(pushStack); i++{
		stack = append(stack,pushStack[i])

		// 当栈不空，且栈头元素相等，弹出当前栈中的元素
		for len(stack) != 0 && stack[len(stack)-1] == popStack[popIndex]{
			stack = stack[:len(stack)-1]
			popIndex++
		}
	}
	return len(stack) == 0
}
