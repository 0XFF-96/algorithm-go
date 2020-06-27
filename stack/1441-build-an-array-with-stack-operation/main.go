package _441_build_an_array_with_stack_operation


func buildArray(target []int, n int) []string {
	// strictly increasing
	// only contains 1 to n
	curIndex := 0
	var stack []int
	ret := []string{}
	for i:=1; i <= n; i++ {
		stack = append(stack, i)
		ret = append(ret, "Push")
		if curIndex > len(target) {
			return ret
		}
		if stack[len(stack)-1] == target[curIndex] {
			curIndex++
			if curIndex == len(target) {
				return ret
			}
		} else {
			ret = append(ret, "Pop")
			stack = stack[:len(stack)-1]
		}
	}
	return ret
}
