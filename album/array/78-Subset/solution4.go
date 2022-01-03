package _8_Subset



func subsetsV4(nums []int) [][]int {
	ans := [][]int{}
	helperV4(nums, 0, []int{}, &ans)
	return ans
}

func helperV4(nums []int, level int, set []int, ans *[][]int) {
	if level == len(nums) {
		*ans = append(*ans, append([]int{}, set...))
		return
	}

	// 这里不用 for 循环？

	// 这里的基本思想是，
	// 取当前元素与不取当前元素之间的区别。
	//
	set = append(set, nums[level])
	helperV4(nums, level+1, set, ans)
	set = set[:len(set)-1]
	helperV4(nums, level+1, set, ans)
}
