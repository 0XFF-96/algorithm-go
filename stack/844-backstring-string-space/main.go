package _44_backstring_string_space


func backspaceCompare(S string, T string) bool {
	var ret1 []rune
	var ret2 []rune

	//
	// 可以选择从后往前遍历
	// 遇到 # 就跳过两个空格
	// done !
	for _, s := range S {
		switch rune(s) {
		case rune('#'):
			if len(ret1) == 0 {
				continue
			}
			ret1 = ret1[:len(ret1)-1]
		default:
			ret1 = append(ret1, rune(s))
		}
	}

	// we could refactor this two block into the same logic
	for _, t := range T {
		switch rune(t) {
		case rune('#'):
			if len(ret2) == 0 {
				continue
			}
			ret2 = ret2[:len(ret2)-1]
		default:
			ret2 = append(ret2, rune(t))
		}
	}

	return string(ret1) == string(ret2)
}
