package array

func sumZero(n int) []int {
	if n == 1 {
		return []int{0}
	}
	// A[i] = i * 2 - n + 1
	ret := make([]int, n)
	for i:=0;i < n;i++{
		ret[i] =  i*2 -n + 1
	}
	return ret
}
