package tree



func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left_correct := ( root.Left == nil || root.Val == root.Left.Val && isUnivalTree(root.Left))
	right_correct := ( root.Right == nil || root.Val == root.Right.Val && isUnivalTree(root.Right))
	return left_correct && right_correct
}

