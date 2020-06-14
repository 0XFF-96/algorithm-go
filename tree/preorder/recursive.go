package preorder


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
