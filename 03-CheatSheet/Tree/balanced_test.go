package tree

import (
	"math"
	"testing"
)

func TestIsBalanced(t *testing.T){
	tree := initTree()
	tree.Left.Left = &TreeNode{Val:1}
	tree.Left.Left.Left = &TreeNode{Val:1}
	b := isBalanced(tree)
	t.Log(b)
}

func isBalanced(root *TreeNode) bool {
	if root ==  nil {
		return true
	}
	lh := height(root.Left)
	rh := height(root.Right)

	if math.Abs(float64(lh - rh )) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

// https://leetcode.com/problems/balanced-binary-tree/discuss/35732/Golang-DFS-solution-(bottom-up)

func isBalancedV2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalanced(root.Left) && isBalanced(root.Right) && abs(depthOfTree(root.Left), depthOfTree(root.Right)) < 2
}



func depthOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depthOfTree(root.Left), depthOfTree(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}