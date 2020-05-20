package _0_valid_parentheses





func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	m := map[byte]byte{
		')':'(',
		'}':'{',
		']':'[',
	}

	// 可以减少类型转换的
	stack := []string{}
	for i:=0;i < len(s);i ++{
		char := s[i]
		v, ok := m[char]
		if ok {
			// we may have edge case when ) was push into stack before any open parathesis
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if string(v) == top {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		} else {
			// if we have open paratheses pop it in to stack
			stack = append(stack, string(char))
		}
	}
	return len(stack) == 0
}
