package array

func checkIfExist(arr []int) bool {
	m := map[int]bool{}

	for _, v := range arr {
		tmp := v*2
		tmp2 := v / 2

		_, ok1 := m[tmp]
		_, ok2 := m[tmp2]
		if ok1 || (ok2 && v %2 == 0) {
			return true
		}
		m[v] = true
	}
	return false
}