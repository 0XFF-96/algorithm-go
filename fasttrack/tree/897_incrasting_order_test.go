package tree


// 看看非递归的答案
//


// 超时了
// ...
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	dummy := &TreeNode{Val:0}
	pre := dummy
	cur := root
	stack := []*TreeNode{}

	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pre.Right = cur
			cur = cur.Right
			pre.Left = nil
			pre = pre.Right
			continue
		}

		stack = append(stack, cur)
		cur = cur.Left
	}
	return dummy.Right
}


// AC 了
func increasingBSTWithRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	ans := &TreeNode{Val:0}
	cur := ans
	var inorder func(root *TreeNode)

	inorder = func(root *TreeNode){
		if root == nil {
			return
		}
		inorder(root.Left)
		root.Left = nil
		cur.Right = root
		cur = root
		inorder(root.Right)
	}
	inorder(root)
	return ans.Right
}

