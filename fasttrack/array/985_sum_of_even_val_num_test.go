package array

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	if len(A) == 0 {
		return nil
	}
	var evenValSum int
	for _, v := range A {
		// bit
		if v % 2 == 0 {
			evenValSum += v
		}
	}

	var ret []int
	// 什么时候 count
	for _, q := range queries {
		idx := q[1]
		val := q[0]
		if A[idx] % 2 == 0 {
			evenValSum -= A[idx]
		}
		A[idx] += val
		if A[idx] % 2 == 0 {
			evenValSum += A[idx]
		}
		ret = append(ret, evenValSum)
	}
	return ret
}
