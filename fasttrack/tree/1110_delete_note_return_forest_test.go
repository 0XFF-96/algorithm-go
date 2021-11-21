package tree

// 能通过，但是不 AC 的答案
// 有一种情况没有 AC, 为什么？
// 请讲出理由？
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root== nil || len(to_delete) == 0 {
		return nil
	}

	var ret []*TreeNode
	var helper func(root *TreeNode) *TreeNode


	helper = func(root *TreeNode)*TreeNode{
		if root == nil {
			return nil
		}

		root.Left = helper(root.Left)
		root.Right = helper(root.Right)

		isDelete := false
		for _, v := range to_delete {
			if v == root.Val {
				isDelete = true
			}
		}

		if isDelete {
			if root.Left != nil {
				ret = append(ret,root.Left)
			}
			if root.Right != nil {
				ret = append(ret, root.Right)
			}
			return nil
		}
		return root
	}

	helper(root)
	return ret
}



// 看一下两个标准答案
// 1、https://leetcode.com/problems/delete-nodes-and-return-forest/discuss/328854/Python-Recursion-with-explanation-Question-seen-in-a-2016-interview
// 2、golang 解法 https://leetcode.com/problems/delete-nodes-and-return-forest/discuss/444521/golang-%3A-90-faster-100-less-user-BFS-and-HashTable
// 3、c++ 方法：https://leetcode.com/problems/delete-nodes-and-return-forest/discuss/444521/golang-%3A-90-faster-100-less-user-BFS-and-HashTable

// 看答案写的
// 有几个 case, 1、delete root 2、delete middle.
func delNodesV2(root *TreeNode, to_delete []int) []*TreeNode {
	newParents := []*TreeNode{}
	toDelete := to_delete

	var delNodesClosure func(*TreeNode, bool)

	delNodesClosure = func(root *TreeNode, isRoot bool){
		if root == nil {
			return
		}
		if !shouldDelete(root, toDelete) && isRoot {
			newParents = append(newParents, root)
		}
		if shouldDelete(root, toDelete) {
			delNodesClosure(root.Left, true)
			delNodesClosure(root.Right, true)
		} else {
			delNodesClosure(root.Left, false)
			delNodesClosure(root.Right, false)
		}
		if shouldDelete(root.Left, toDelete) {
			root.Left = nil
		}
		if shouldDelete(root.Right, toDelete) {
			root.Right = nil
		}
	}
	delNodesClosure(root, true)

	return newParents
}

func shouldDelete(node *TreeNode, toDelete []int) bool {
	if node == nil {
		return true
	}
	for _, toDel := range toDelete {
		if node.Val == toDel {
			return true
		}
	}
	return false
}