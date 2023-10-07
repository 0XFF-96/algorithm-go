package merge_sort

// 从尾部打印链表
type LinkNode struct {
	val int
	next *LinkNode
}

func mergeSortedLinkList(link1 * LinkNode, link2 *LinkNode) *LinkNode{
	if link1 == nil {
		return link2
	}
	if link2 == nil {
		return link1
	}

	// 这个逻辑可以简化一下
	// 用一个 for 循环就 OK 了
	//
	// 如何设置 dummy 节点
	result := &LinkNode{}
	dummy := result
	for ; link1 != nil && link2 != nil; {
		if link1.val < link2.val {
			dummy.val = link1.val
			link1 = link1.next
		} else {
			dummy.val = link2.val
			link2 = link2.next
		}
		dummy.next = &LinkNode{}
		dummy = dummy.next
	}

	// link1 left
	for ; link1 != nil ;{
		dummy.val = link1.val
		link1 = link1.next
	}

	// link2 left
	for ; link2 != nil ;{
		dummy.val = link2.val
		link2 = link2.next
	}
	return result
}
