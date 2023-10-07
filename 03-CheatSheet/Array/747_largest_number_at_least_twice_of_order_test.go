package array


func dominantIndex(nums []int) int {
	largestIdx := 0
	num := 0
	for idx, v := range nums {
		if v >= num {
			num = v
			largestIdx = idx
		}
	}

	for idx, v := range nums {
		if num < 2 * v && idx != largestIdx {
			return -1
		}
	}
	return largestIdx
}

// follow-up
// 有没有办法能否避免第二次循环呢？


// 题目转化为: second largest idx
// https://leetcode.com/problems/largest-number-at-least-twice-of-others/solution/
// 只需要和第二大的元素比较一次，即可。
// 空间复杂度呢？

