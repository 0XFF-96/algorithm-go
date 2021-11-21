package tree

import (
	"testing"
)

// 从中序遍历序列出发...
// [-10, -3, 0, 5, 9]

// BST 有什么性质？
// 递归
func TestSortArryToBST(t *testing.T){
	tree := sortedArrayToBST([]int{2, 1, 3})
	t.Log(tree)
	t.Log(tree.Left.Val)
	t.Log(tree.Right.Val)
}

func sortedArrayToBST(nums []int)*TreeNode{
	if len(nums) == 0 {
		return nil
	}
	length := len(nums)

	// 下次有这种除法的地方都要小心...
	rootIndex := length / 2
	root := &TreeNode{Val:nums[rootIndex]}
	root.Left = sortedArrayToBST(nums[:rootIndex])
	root.Right = sortedArrayToBST(nums[rootIndex+1:])
	return root
}

// func sortedArrayToBST()



// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/discuss/35281/Golang-solution-using-backtracking-(DFS)
// 递归解法