func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}

	res := make([]int, rowIndex+1)
	res[0], res[1] = 1, 1
	for i := 2; i <= rowIndex; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 || j == i {
				res[j] = 1
				continue
			}

			res[j] = res[j] + res[j-1]
		}
	}
	return res
}