package main

type ListNode struct {
	    Val int
	     Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// return the middle of the second
	// if fast != nil {
	//     slow = slow.Next
	// }
	return slow
}
