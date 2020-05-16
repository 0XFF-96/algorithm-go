package main

type ListNode struct {
	     Val int
	     Next *ListNode
}

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	// 把值搬上来，然后最后一个节点释放调就可以了
	// 冗余了
	// 不需要 for 循环♻️
	for node != nil && node.Next != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
			break
		}
		node = node.Next
	}
}

func deleteNodeV2(node *ListNode) {
	if node == nil {
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

