

// 使用降维的思维模式
// 
func removeKdigits(num string, k int) string {
	res := []rune{}

	for _, v := range num {
		for len(res) > 0 && res[len(res)-1] > v && k > 0 {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, v)
	}

	// edge case
	for k > 0 {
		k--
		res = res[:len(res)-1]
	}

	// trim leading zeros
	for len(res) > 0 && res[0] == '0' {
		res = res[1:len(res)]
	}

	if len(res) == 0 {
		return "0"
	}
	return string(res)
}