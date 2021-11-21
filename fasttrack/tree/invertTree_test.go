package tree


// 递归解法
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Right = left
	root.Left = right
	return root
}

// 有没有