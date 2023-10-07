package recursive

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	inordermap := make(map[int]int)
	for i, in := range inorder {
		inordermap[in] = i
	}
	return helper(preorder, 0, len(preorder) - 1, inorder, 0, len(inorder) - 1, inordermap)
}

func helper(preorder []int, preLeft, preRight int, inorder []int, inLeft, inRight int, inordermap map[int]int) *TreeNode {
	if preLeft > preRight {
		return nil
	}

	rootVal := preorder[preLeft]
	root := &TreeNode{Val: rootVal}

	rootPos := inordermap[rootVal]
	leftLen := (rootPos - 1) - inLeft + 1

	root.Left = helper(preorder, preLeft + 1, preLeft + leftLen, inorder, inLeft, rootPos - 1, inordermap)
	root.Right = helper(preorder, preLeft + leftLen + 1, preRight, inorder, rootPos + 1, inRight, inordermap)
	return root
}
