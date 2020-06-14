package tmp

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

// 题目框架
// You may assume that duplicates do not exist in the tree.
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 递归
	// 先找到根， 然后。左右递归

	// root := find(root)
	// root.Left =   (preorder, inorder)
	// root.Right =  (preorder, inorder)

	// 非递归怎么做？
	// 利用栈？
	preEnd := len(preorder)
	inEnd := len(inorder)
	return buildInOrderTree(preorder, inorder, 0, preEnd, 0, inEnd)
}

// 关键是：找出这几个坐标之间的关系
//
func buildInOrderTree(preorder []int, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	// 找到这个未知
	// for
	// 要么用 map
	// root.Left =
		// root.Right =

	return root
}
