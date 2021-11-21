package array




// 代码写得太难看了
// 有没有优化的方法？
func plusOne(digits []int) []int {
	skip := false
	carry := 0
	tmp := 0
	for i:=len(digits)-1;i>=0;i-- {
		if skip {
			break
		}
		if i == len(digits) - 1 {
			tmp = digits[i] + 1
		} else {
			tmp = digits[i] + carry
		}
		if tmp >= 10 {
			carry = 1
			// got := tmp % 10
			digits[i] = tmp % 10
			continue
		}
		digits[i] = tmp
		carry = 0
		skip = true
	}

	if carry == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
