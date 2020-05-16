package main

type ListNode struct {
	     Val int
	     Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast := head
	slow := head
	stack := []int{}
	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 单数队列
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != stack[len(stack)-1] { return false }
		slow = slow.Next
		stack = stack[:len(stack)-1]
	}
	return true
}
