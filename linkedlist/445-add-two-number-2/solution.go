package main

type ListNode struct {
	     Val int
	     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//  use extrat space
	// 对齐？
	// 递归？
	// 先递归深入
	// 不能反转链表？
	// leading zero 开头没有零值。
	// 对齐才可以？
	// 用栈

	l1Stack := []*ListNode{}
	l2Stack := []*ListNode{}

	curL1 := l1
	for curL1 != nil {
		l1Stack = append(l1Stack, curL1)
		curL1 = curL1.Next
	}

	curL2 := l2
	for curL2 != nil {
		l2Stack = append(l2Stack, curL2)
		curL2 = curL2.Next
	}

	dummy := &ListNode{}
	curDummy := dummy

	var carry int
	for len(l1Stack) != 0 || len(l2Stack) != 0 {
		var topL1 int
		var topL2 int
		if len(l1Stack) != 0 {
			topL1 = l1Stack[len(l1Stack)-1].Val
			l1Stack = l1Stack[:len(l1Stack)-1]
		}

		if len(l2Stack) != 0 {
			topL2 = l2Stack[len(l2Stack)-1].Val
			l2Stack = l2Stack[:len(l2Stack)-1]
		}
		num := topL1 + topL2 + carry

		node := &ListNode{
			Val: num % 10,
		}

		carry = num / 10
		// 头插法
		node.Next = curDummy.Next
		curDummy.Next = node
	}
	if carry == 1 {
		node := &ListNode{
			Val: 1,
		}

		node.Next = curDummy.Next
		curDummy.Next = node
	}

	return dummy.Next
}
