package tree

import (
	"testing"
)


// TODO: 参考一下非递归解法
// https://leetcode.com/problems/sum-of-left-leaves/discuss/88950/Java-iterative-and-recursive-solutions

func TestSumOfLeftLeaves(t *testing.T){


}

func sumOfLeftLeavesNoRecursive(root *TreeNode) int {
	if root == nil {return 0}
	sum := 0


	// 中序遍历框架的转变



	return sum
}

// AC 的答案...
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0

	var  helper = func(root *TreeNode, isLeft bool) {}

	helper = func(root *TreeNode, isLeft bool){
		if root == nil {
			return
		}

		if isLeft && root.Left == nil && root.Right == nil {
			sum += root.Val
		}

		helper(root.Left, true)
		helper(root.Right, false)
	}

	helper(root.Left, true)
	helper(root.Right, false)
	return sum
}


func sumOfLeftLeavesWrong(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0

	var  helper = func(root *TreeNode, isLeft bool) {}

	helper = func(root *TreeNode, isLeft bool){
		if root == nil {
			return
		}

		if isLeft {
			sum += root.Val
		}

		helper(root.Left, true)
		helper(root.Right, false)
	}

	return sum
}