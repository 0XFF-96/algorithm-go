package main 


// 利用的是重排相关数字的相关技巧。 
// 
func maximumSwap(num int) int {
	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	var maxDigit, maxIndex int
	index1, index2 := -1, -1
	for i, digit := range digits {
		if digit > maxDigit {
			maxDigit = digit
			maxIndex = i
		} else if digit < maxDigit {
			index1 = i
			index2 = maxIndex
		}
	}

	if index1 >= 0 {
		digits[index1], digits[index2] = digits[index2], digits[index1]
	}

	var res int
	for i := len(digits) - 1; i >= 0; i-- {
		res = res*10 + digits[i]
	}
	return res
}