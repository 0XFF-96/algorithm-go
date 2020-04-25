package main 

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		// 这里会溢出吗？
		// 和二分查找不同
		m := (l + r) >> 1
		if target < nums[m] {
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		} else {
			// 等于就直接返回？		
			// 有多个重复元素的时候呢？
			return m
		}
	}
	return l
}