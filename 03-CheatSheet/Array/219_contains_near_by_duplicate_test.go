package array


// 程序运行的效率不高
// 看看需要怎么提高效率...
//
func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}

	for idx, v := range nums {
		n, ok := m[v]
		if ok && (idx - n) <= k{
			return true
		}
		m[v] = idx
	}
	return false
}


