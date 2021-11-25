package linklist


//
// [4,5,1,9]
// 5
//
// [4,1,9]
// The given node will not be the tail and it will always be a valid node of the linked list.
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	// 把值搬上来，然后最后一个节点释放调就可以了
	for node != nil && node.Next != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
			break
		}
		node = node.Next
	}
}

// 根据相关限制条件.
func deleteNodeV2(node *ListNode) {
	node.Val=node.Next.Val
	node.Next=node.Next.Next
}


