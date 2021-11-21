package linklist

// A linked list can be reversed either iteratively or recursively.
// Could you implement both?

func reverseList(head *ListNode) *ListNode {
	// reverseList
	//
	if head == nil {
		return nil
	}
	var prev *ListNode
	cur := head
	// var next *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}


// 递归方法
//
// 分析一下栈调用
// 为什么可以这样做？
// 是错的？
func reverseListV2(head *ListNode) *ListNode {
	// reverseList
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseListV2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
