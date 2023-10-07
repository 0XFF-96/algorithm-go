package tree



func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	left := mergeTrees(t1.Left, t2.Left)
	right := mergeTrees(t1.Right, t2.Right)

	return &TreeNode{
		Val: t1.Val + t2.Val,
		Left: left,
		Right: right,
	}
}



// follow-up
// 有没有非递归解法？
// 有没有
