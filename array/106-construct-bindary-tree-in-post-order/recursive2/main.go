package recursive2


// 没有 AC
// 有一点错误❌
func buildTree(inorder []int, postorder []int) *TreeNode {
	return helper(inorder, postorder, 0, len(inorder)-1, len(postorder)-1)
}

func helper(inorder []int, postorder []int, inStart, inEnd ,lastIndex int) *TreeNode {
	if inStart > inEnd || lastIndex < 0 {
		return nil
	}
	rootVal := postorder[lastIndex]
	root := &TreeNode{
		Val:rootVal,
	}
	var inIndex int
	for i:=inStart; i <= inEnd; i++{
		if inorder[i] == rootVal {
			inIndex = i
		}
	}

	root.Left = helper(inorder, postorder, inStart, inIndex -1, lastIndex - inIndex - inEnd +1)
	root.Right = helper(inorder, postorder,  inIndex+1, inEnd, lastIndex -1)
	return root
}
