package tree


// 前序遍历的递归于非递归算法
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// preorderTraversal()
	stack := []*TreeNode{}
	cur := root
	var ret []int
	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			cur = cur.Right
			continue
		}


		ret = append(ret, cur.Val)
		stack = append(stack, cur)
		cur = cur.Left
	}
	return ret
}
// 中: 左根右
// 前：根左右
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var st []*TreeNode
	st = append(st, root)
	for len(st)!=0 {
		l := len(st)
		res = append(res, st[l-1].Val)
		curr := st[l-1]
		st = st[:l-1]
		if curr.Right!= nil {
			st = append(st,curr.Right)
		}
		if curr.Left!=nil {
			st = append(st,curr.Left)
		}
	}
	return res
}




// 值得学习的地方
// 1、学习如何使用 reference 在这个例子里面
func preorderTraversalWithRecursive(root *TreeNode) []int {
	var result []int
	preorder(root,&result)
	return result
}
func preorder(root *TreeNode,output *[]int) {
	if root != nil {
		*output = append(*output,root.Val)
		preorder(root.Left,output)
		preorder(root.Right,output)
	}
}
