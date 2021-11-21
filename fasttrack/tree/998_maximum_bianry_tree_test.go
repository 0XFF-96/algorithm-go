package tree




// 为什么只需要 find the solution of right sub tree ?
// 不懂...
// 为什么不同插入左子树呢？
// 算法框架
// 这题完全想错了


// Maximum binary tree
// https://www.youtube.com/watch?v=oHnT9zTsXTg

//
//func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
//	if root == nil {
//		return &TreeNode{Val:val}
//	}
//
//	if val > root.Val {
//		node := &TreeNode{Val:val}
//		node.Left = root
//		return node
//	}
//
//	if val > root.Left && val < root.Right {
//		return insertIntoMaxTree(root.Right, val)
//	}
//
//	if val > root.Right && < root.Left {
//		return insertIntoMaxTree(root.Left, val)
//	}
//
//	if val > root.Right && root.Right > root.Left {
//
//	}
//
//	if val > root.Left && root.Left > root.Right {
//
//	}
//
//	return
//}


// https://leetcode.com/problems/maximum-binary-tree-ii/discuss/375129/Go-0ms-Recursive-Solution
//


// 他的题目含义是不是包括，总是数组的右子树插入节点？
// V1 解法
func insertIntoMaxTreeV1(root *TreeNode, val int) *TreeNode {
	if root != nil && root.Val > val {
		root.Right = insertIntoMaxTreeV1(root.Right, val)
		return root
	}
	node := &TreeNode{Val: val, Left: root}
	node.Left = root
	return node
}