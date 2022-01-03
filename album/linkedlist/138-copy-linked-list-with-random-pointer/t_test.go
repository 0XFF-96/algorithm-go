package main



func copyRandomListV2(head *Node) *Node {
	if head == nil {
		return nil
	}

	// copy 链表
	cur := head
	for cur != nil {
		copyNode := &Node{Val:cur.Val}
		next := cur.Next
		cur.Next = copyNode
		copyNode.Next = next

		cur = copyNode.Next
	}

	// copy 链表
	cur = head
	for cur != nil && cur.Next != nil {
		random := cur.Random
		if random == nil {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = random.Next
		}
		cur = cur.Next.Next
	}

	// split it into two linked list
	// must do it this way
	headPrev := &Node{}
	retPrev := &Node{}
	prev := retPrev
	for t := head; t != nil; {
		headPrev.Next = t
		prev.Next = t.Next
		headPrev = headPrev.Next
		prev = prev.Next
		t = t.Next.Next
	}
	headPrev.Next = nil
	prev.Next = nil

	return retPrev.Next

	// 这是什么错误？
	// Next pointer of node with label 7 from the original list was modified.
}