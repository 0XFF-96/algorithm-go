package recursive2


///**
// * Definition for a binary tree node.
//
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

// You may assume that duplicates do not exist in the tree.
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	return helper(0, 0, len(inorder)-1, preorder, inorder)
}


// inEnd 是包含关系
func helper(preStart, inStart, inEnd int, preorder []int, inorder []int)*TreeNode{
	if preStart > len(preorder) -1 || inStart > inEnd {
		return nil
	}
	root := &TreeNode{
		Val:preorder[preStart],
	}
	var inIndex int

	for i:=inStart;i <=inEnd; i++{
		if inorder[i] == root.Val {
			inIndex = i
		}
	}

	root.Left = helper(preStart+1, inStart, inIndex-1, preorder, inorder)

	root.Right = helper(
		preStart + inIndex - inStart +1,
		inIndex +1,
		inEnd,
		preorder,
		inorder,
	)
	return root
}
