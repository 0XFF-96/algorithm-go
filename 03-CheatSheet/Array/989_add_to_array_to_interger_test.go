package array


// 错误解法，
// 超出数字限制。
func addToArrayForm(A []int, K int) []int {
	sum := A[0]
	for i:=1;i < len(A);i++ {
		sum = sum * 10
		sum += A[i]
	}
	// 超过数字限制...
	sum += K
	var ret []int
	if sum == 0 {
		return []int{0}
	}
	for sum > 0 {
		ret = append(ret, sum % 10)
		sum /= 10
	}

	// reverse
	l := 0
	j := len(ret) -1
	for l < j {
		ret[l],ret[j] = ret[j], ret[l]
		l++
		j--
	}

	return ret
}


// 分析为什么会错？
func addToArrayFormV2(A []int, K int) []int {
	var kA []int
	for K > 0 {
		kA = append(kA, K%10)
		K /= 10
	}
	reverseV(kA)
	carry := 0
	var ret []int
	for len(kA) > 0 && len(A) > 0 {
		numK := kA[len(kA)-1]
		numA := A[len(A)-1]

		sum := numK + numA + carry
		carry = sum / 10
		sum = sum % 10

		ret = append(ret, sum)
		kA = kA[:len(kA)-1]
		A = A[:len(A)-1]
	}

	if len(kA) > 0 {
		ret = append(ret, kA[len(kA)-1]+carry)
		ret = append(ret, kA[:len(kA)-1]...)
		carry = 0
	}

	if len(A) > 0 {
		ret = append(ret, A[len(A)-1]+carry)
		ret = append(ret, A[:len(A)-1]...)
		carry = 0
	}
	if carry == 1 {
		ret = append(ret, 1)
	}


	return reverseV(ret)
}


// 非长简洁版本
func addToArrayFormV3(A []int, K int) []int {

	i := len(A)-1
	cur := K
	var ret []int
	for i >=0 || cur >0 {
		if i >= 0 {
			cur += A[i]
		}
		ret = append(ret, cur % 10)
		cur /= 10
		i--
	}

	return reverseV(ret)
}


func reverseV(ret []int)[]int{
	l := 0
	r := len(ret) -1
	for l <= r {
		ret[l], ret[r] = ret[r], ret[l]
		l++
		r--
	}
	return ret
}
