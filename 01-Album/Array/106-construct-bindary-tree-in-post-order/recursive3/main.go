package recursive3


// 这个更加简洁
func buildTree(inorder []int, postorder []int,) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Find the root value's index in the inorder array
	var idx int
	for k, v :=  range inorder {
		if v == postorder[len(postorder) - 1] {
			idx = k
			break
		}
	}

	// This is how to devide the original array by the index we get
	// post_left, post_right := postorder[0: idx], postorder[idx: -1]
	// in_left, in_right := inorder[0: idx], inorder[idx+1:]

	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(inorder[0: idx], postorder[0: idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1]),
	}
}

