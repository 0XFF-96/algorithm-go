package array


// 相关的事情。
//
func sortArrayByParity(A []int) []int {
	if len(A) == 0 {
		return nil
	}
	// two points
	i :=0
	j := len(A)-1

	for i < j {
		if A[i] % 2 == 0 {
			i++
		} else if A[i] %2 != 0 && A[j] %2 == 0 {
			A[i],A[j] = A[j],A[i]
			i++
			j--
		} else {
			j--
		}
	}
	return A
}


