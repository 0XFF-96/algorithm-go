package array

// https://www.youtube.com/watch?v=y8otfJ-5X40
//
//

// 记住这个规则。
// If X is the first i digits of the array as a binary number,
// then 2X + A[i] is the first i+1 digits.
//


func prefixesDivBy5(A []int) []bool {
	// 怎么 trasfer string to bit ?
	// 移位？
	ret := make([]bool, len(A))
	prefixSum := make([]int, len(A))

	if A[0] == 0 {
		ret[0] = true
		prefixSum[0] = 0
	} else {
		ret[0] = false
		prefixSum[0] = 1
	}

	var start int
	for i:=1; i < len(A); i ++ {
		start = prefixSum[i-1]*2
		if A[i] == 1 {
			start = start + 1
		}
		if start % 5 == 0 {
			ret[i] = true
		} else {
			ret[i] = false
		}
		// 防止输入，数据过大，造成溢出
		//
		// 所有需要 convert 数据的地方，都需要注意这一点
		//
		start = start % 5
		prefixSum[i] = start
	}
	return ret
}
