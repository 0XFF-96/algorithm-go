package array



func findSpecialInteger(arr []int) int {
	l := len(arr)

	count := 1
	for i:=1; i < len(arr);i++ {
		if arr[i-1] == arr[i] {
			count++
		} else {
			count = 1
		}
		if float64(count) > float64(l / 4) {
			return arr[i]
		}
	}

	return arr[len(arr)-1]
}
