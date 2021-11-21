package archive_Tree

import (
	"math"

	"pratice-go/algorithmTopic/tree"
)

// review-time: 1,
// 题解：https://leetcode.com/problems/validate-binary-search-tree/solution/
// 有三种方法可以解决这道题目：1、递归 2、中序遍历 3、Morri 的方法
// 相关情况。
// 递归解法是:
// 中序遍历一定是有序特性
// 考点是 BST 的性质，能否转化问题。
//
func isValidBST(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	cur := root
	stack := []*tree.TreeNode{}
	preVal := math.MinInt64
	for len(stack)!= 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if cur.Val > preVal {
				preVal = cur.Val
			} else {
				return false
			}
			cur = cur.Right
			continue
		}

		stack = append(stack, cur)
		cur = cur.Left
	}
	return true
}

// Morii 的方法构建临时边
func isValidBSTMorii(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	pre := math.MinInt64
	cur := root
	for cur != nil {
		if cur.Left == nil {
			if pre >= cur.Val {
				return false
			}
			pre = cur.Val
			cur = cur.Right
		} else {
			preNode := cur.Left
			for preNode.Right != nil && preNode.Right != cur {
				preNode = preNode.Right
			}

			if preNode.Right == nil {
				// 创造临时边
				preNode.Right = cur
				cur = cur.Left
			} else {
				// 临时边
				preNode.Right = nil
				if pre >= cur.Val {
					return false
				} else {
					pre = cur.Val
				}
				cur = cur.Right
			}
		}
	}
	return true
}