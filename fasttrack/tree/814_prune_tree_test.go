package tree


func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := pruneTree(root.Left)
	right := pruneTree(root.Right)

	if root.Val != 1 && left == nil && right == nil{
		return nil
	}

	root.Left = left
	root.Right = right

	return root
}

