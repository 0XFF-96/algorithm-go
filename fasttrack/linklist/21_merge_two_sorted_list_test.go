package linklist



func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{Val:0}
	cur := dummyNode
	for l1 != nil && l2 != nil {
		head1 := l1.Val
		head2 := l2.Val
		if head1 <= head2 {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}
	return dummyNode.Next
}
