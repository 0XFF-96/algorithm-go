package linklist

// 这题想了这么久，不应该啊
// 从当前状态，与当前的下一状态进行比较
// 一开始选取的比较的基点，就不正确🙆。
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}


// 错误代码示例
func deleteDuplicatesWrong(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		if pre == nil  {
			pre = cur
			cur = cur.Next
		}
		if pre.Val == cur.Val {
			cur.Next = cur.Next.Next
		}

	}
	return head
}