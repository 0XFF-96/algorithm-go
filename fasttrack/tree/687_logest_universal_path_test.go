package tree

//
// 二叉树中有一种解法：This path may or may not pass through the root.
// 这题都是用递归做，而且都在返回都时候，有一个特点。
// 
// 一定要联想到
// 如果求一颗树的所有 path 怎么求？
// path sum


// 相关题目：
// 1、path sum
// 2、Count Univalue Subtrees. subtrees 的结构...
// 3、Binary Tree Maximum Path Sum...=


// 有一个 case 没有考虑到？
// [5,4,5,1,1,5] pass
// [5,4,5,1,1,5] not pass


// 错误的地方
// 没有考虑到左右节点都存在都情况， 提早 return 了
//
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := longestUnivaluePath(root.Left)
	right := longestUnivaluePath(root.Right)
	rowLeft, rowRight := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {

		// 这里写错了
		// return left + 1
		rowLeft += left + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {

		// 这里写错了
		// return right +1
		rowRight += right + 1
	}

	if rowLeft > rowRight {
		return  rowLeft
	} else {
		return rowRight
	}
}

func longestUnivaluePathAC(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := -1
	var helper func(root *TreeNode) int

	helper = func(root *TreeNode)int {
		if root == nil {
			return 0
		}

		left := helper(root.Left)
		right := helper(root.Right)
		arrowLeft := 0
		arrowRight := 0
		if root.Left != nil && root.Left.Val == root.Val {
			arrowLeft += left + 1
		}
		if root.Right != nil && root.Right.Val == root.Val {
			arrowRight += right + 1
		}
		ans = max(ans, arrowLeft + arrowRight)
		return max(arrowLeft, arrowRight)
	}
	helper(root)
	return ans
}


// 复杂度分析
// Complexity Analysis

// Time Complexity: O(N)O(N), where NN is the number of nodes in the tree. We process every node once.
// Space Complexity: O(H)O(H), where HH is the height of the tree. Our recursive call stack could be up to HH layers deep.
// https://leetcode.com/problems/longest-univalue-path/discuss/167989/Golang-bottom-up-recursion-solution
// 什么是 tree 的 bottom-up solution ?
// 尝试解一下...
//