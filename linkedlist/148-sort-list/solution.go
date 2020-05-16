package main

type ListNode struct {
	     Val int
	     Next *ListNode
}

// 应该还有其他几种不同的做法的
//
func sortList(head *ListNode) *ListNode {
	// bottom-up
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// ?
	prev.Next = nil

	l1 := sortList(head)
	l2 := sortList(slow)

	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	l := &ListNode{}
	p := l

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}

	return l.Next
}

