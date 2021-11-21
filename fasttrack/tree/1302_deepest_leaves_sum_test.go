package tree

func deepestLeavesSum(root *TreeNode) int {
	m := map[int]int{}
	maxDepth := 0

	// 递归找到最深的节点
	var helper func(root *TreeNode, depth int)
	helper = func(root *TreeNode, depth int){
		if root == nil {
			return
		}
		depth++
		if depth > maxDepth {
			maxDepth = depth
		}
		m[depth] = m[depth] + root.Val
		helper(root.Left, depth+1)
		helper(root.Right, depth+1)
	}

	helper(root, 0)
	return m[maxDepth]
}

