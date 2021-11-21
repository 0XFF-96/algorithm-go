package tree

import "math"

// 成就感
// 一次过 AC

// https://leetcode.com/problems/maximum-binary-tree/solution/
// 有相关 solution 的题目
//
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// find the max value index of a given array
	index := findMaxValueIndex(nums)
	left := constructMaximumBinaryTree(nums[:index])
	right := constructMaximumBinaryTree(nums[index+1:])
	root := &TreeNode{Val:nums[index]}
	root.Left = left
	root.Right = right
	return root
}


func findMaxValueIndex(nums[]int) int {
	if len(nums) == 0 {
		return -1
	}
	maxIndex := -1
	maxValue := math.MinInt32
	for idx, n := range nums {
		if n > maxValue {
			maxIndex = idx
			maxValue  = n
		}
	}
	return maxIndex
}