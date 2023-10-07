
// 递归的方法
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

	// 这个 p 是否在程序运行过程中一直没有发生过变化？
	// ？
	p := reverseListV2(head.Next)

	// 这代表什么
	head.Next.Next = head
	head.Next = nil
	return p
}
