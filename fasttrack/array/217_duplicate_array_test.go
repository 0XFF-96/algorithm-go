package array

// 这是一道简单题
// 应该关注性能层面的提高
// 应该关注效率和时间空间复杂度的提高。

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}

	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}
// sorted it
// and two pointer



