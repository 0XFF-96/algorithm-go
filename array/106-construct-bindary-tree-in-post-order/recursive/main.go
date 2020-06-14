package recursive

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return helper(inorder, postorder, 0, len(inorder)-1, len(postorder)-1)
}

func helper(inorder []int, postorder []int, inStart, inEnd ,lastIndex int) *TreeNode {
	if inStart > inEnd {
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

	root.Left = helper(inorder, postorder, inStart, inEnd - inIndex -1, lastIndex - inIndex +1)
	root.Right = helper(inorder, postorder,  inIndex+1, inEnd, lastIndex -1)
	return root
}
