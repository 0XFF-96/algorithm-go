package tree

// 看答案一次 AC
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	var helper func(root *TreeNode) *TreeNode

	helper = func(root *TreeNode) *TreeNode {
		switch {
		case root == nil:
			return nil
		case root.Val > R:
			return helper(root.Left)
		case root.Val < L:
			return helper(root.Right)
		default:
			root.Left = helper(root.Left)
			root.Right = helper(root.Right)
			return root
		}
	}
	return helper(root)
}