package main

// 很厉害的题解
// 结合了 early return 的思想在里面。
func plusOne(digits []int) []int {
    var n int = len(digits)
    for i:= n-1; i>=0; i--{
        if digits[i] < 9 {
            digits[i]+=1
            return digits
        } else {
            digits[i] = 0
        }
    }
    var a = make([]int,n+1)
    a[0] = 1
    return a
}

// 普通的解法？
func plusOneV2(digits []int) []int {
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