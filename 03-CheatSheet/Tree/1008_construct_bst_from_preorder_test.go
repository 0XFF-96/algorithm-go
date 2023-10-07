package tree


func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val:preorder[0]}
	}

	rootVal := preorder[0]
	idx := 0
	for i:=1;i<len(preorder);i++{
		if preorder[i] > rootVal {
			idx = i
			break
		}
	}

	// ⚠️ index 问题
	root := &TreeNode{Val:rootVal}
	root.Left = bstFromPreorder(preorder[1:idx])
	root.Right = bstFromPreorder(preorder[idx:])
	return root
}


// 没有 right 的情况。
// 枚举所有有可能出现的情况。
func bstFromPreorderV2(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val:preorder[0]}
	}

	rootVal := preorder[0]
	idx := 0
	for i:=1;i<len(preorder);i++{
		if preorder[i] > rootVal {
			idx = i
			break
		}
	}

	// 如果没有左右子树的问题
	// 4 2
	// ⚠️ index 问题
	root := &TreeNode{Val:rootVal}
	// 没有 right tree 的情况
	if idx == 0 {
		root.Left = bstFromPreorder(preorder[1:])
		return root
	}
	root.Left = bstFromPreorder(preorder[1:idx])
	root.Right = bstFromPreorder(preorder[idx:])
	return root
}
