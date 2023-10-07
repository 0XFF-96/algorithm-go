package array


// 看看相关的题目
// https://leetcode.com/problems/move-zeroes/solution/
//
func moveZeroes(nums []int)  {
	if len(nums) <= 1 {
		return
	}

	// 反着来
	// 不等于 0 的才需要加..
	l:=0
	for idx, v := range nums {
		if v != 0 {
			nums[idx], nums[l] = nums[l], nums[idx]
			l++
		}
	}
	return
}

