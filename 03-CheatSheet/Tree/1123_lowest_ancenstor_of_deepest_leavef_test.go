package tree


// lca 问题：
// https://www.geeksforgeeks.org/lowest-common-ancestor-in-a-binary-search-tree/
// 这个问题的思考模型是什么？
// 让我困惑的是，他的输入输入出...

// 首先，你得看懂题目
// 为什么在这种情况下，返回 2， 4， 5
// 从 1， 2， 3， 4， 5，
// 开始分析，case study ....
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	res, _ := helperlca(root, 1)
	return res
}

func helperlca(root *TreeNode, lv int) (*TreeNode, int) {
	if root == nil {
		return nil, lv-1
	}

	left, ll := helperlca(root.Left, lv +1)
	right, lr := helperlca(root.Right, lv + 1)

	// 为什么在这里要 return ll ?
	// 搞不明白
	if ll == lr {
		return root, ll
	}

	if ll > lr {
		return left, ll
	}

	return right, lr
}
