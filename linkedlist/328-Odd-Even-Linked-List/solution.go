package main


type ListNode struct {
	Val int
	Next *ListNode
}
// 读懂题目没有？
// we are talking about the 【node number】
// not 【the value in the nodes 】!
// 读懂这句话
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	//
	// ?
	// ?
	// 1->2->nil  ?
	//  1->2->3->nil ?
	// 限制The relative order inside both the even and odd groups should remain as it was in the input.
	//  The first node is considered odd, the
	// second node even and so on...
	odd.Next = evenHead
	return head
}
