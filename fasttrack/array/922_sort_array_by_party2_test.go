package array



// 这里不够好
// 需要一个优化。
//
func sortArrayByParityII(A []int) []int {
	if len(A) % 2 != 0 {
		return nil
	}

	ret := make([]int, len(A))

	t := 0
	for _, v := range A {
		if v % 2 == 0 {
			ret[t] = v
			t += 2
		}
	}

	t = 1
	for _, v := range A {
		if v % 2 == 1 {
			ret[t] = v
			t+=2
		}
	}
	return ret
}
