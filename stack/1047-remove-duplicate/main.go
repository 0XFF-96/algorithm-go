package _047_remove_duplicate

func removeDuplicates(S string) string {
	// 递归
	// 使用栈
	stack := make([]byte, 0, len(S))
	for i := range S {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}


// string  和 rune 之间的转换
// 之后，需要优化的地方。
func removeDuplicatesV2(S string) string {
	stack := make([]rune, 0)
	for _, c := range S {
		if len(stack) > 0 && stack[len(stack) - 1] == c {
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, c)
		}
	}
	result := ""
	for _, c := range stack {
		result += string(c)
	}
	return result
}