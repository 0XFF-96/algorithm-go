package array

import (
	"sort"
)

// Height Checker
// https://leetcode.com/problems/height-checker/discuss/299265/Not-a-well-defined-problem
// https://leetcode.com/problems/height-checker/discuss/300472/Java-0ms-O(n)-solution-no-need-to-sort
// From the description,
// I feel it's a longest non-decreased subsequence problem.
// But also couldn't understand the example well


// 一旦理解了题目之后，就不难了
// 难的是，如何进行优化
//
func heightChecker(heights []int) int {
	// 逆序对
	// Notice that when a group of students is selected they can reorder in any possible way between themselves
	// 冒泡排序 + 计数
	var origin []int
	origin = append(origin, heights...)
	sort.Ints(heights)
	var count int
	for idx, v := range origin {
		if heights[idx] != v {
			count++
		}
	}
	return count
}