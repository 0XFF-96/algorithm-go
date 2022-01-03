package _09_convert_linked_list_to_tree

type ListNode struct {
	     Val int
	     Next *ListNode
}
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	// 难点，不止有一个答案
	// 中序遍历模板
	// 递归构建树的过程
	// 中序遍历，逆向过程

	// 中序遍历的 stack 使用，逆向过程
	if head == nil {
		return nil
	}

	return helper(head, nil)
}

func helper(start *ListNode, end *ListNode) *TreeNode {
	if start == end {
		return nil
	}

	// find middle value
	slow := start
	fast := start
	for fast != end && fast.Next != end {
		slow = slow.Next
		fast = fast.Next.Next
	}

	node := &TreeNode{Val: slow.Val}

	node.Left = helper(start, slow)
	node.Right = helper(slow.Next, end)
	return node
}
