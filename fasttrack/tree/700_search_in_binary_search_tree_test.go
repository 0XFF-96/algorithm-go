package tree

// AC
// https://leetcode.com/problems/search-in-a-binary-search-tree/
// 有没有其他方法可以做一下？
//
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return root
	}
}


