package linklist


// 这一题，用的模板和 remove_duplicate_from_sorted_list
// 是一样的。
// 一开始的时候，做错了。用错方法了。
// 内存消耗比较高。
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	// 注意点一⚠️
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}


// 这个答案错在哪里了？
//

//Wrong Answer
//Runtime: 0 ms
//Your input
//[1,2,6,3,4,5,6]
//6
//Output
//[1,2,3,4,5,6]
//Expected
//[1,2,3,4,5]
func removeElementsV2(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	curDummy := dummy
	cur := head

	// 应该从 Next 开始判断？
	// 还是从当前节点判断？
	for cur != nil {
		if cur.Val == val {
			// 上面条件出现在最后时，直接跳过了
			// 没有办法加入到 dummy 里面。
			cur = cur.Next
			continue
		}
		// 问题出现在这里没有 else 的情况
		// 为什么最后一个数字，不能被正确加入到数组里面。
		curDummy.Next = cur
		curDummy = curDummy.Next
		cur = cur.Next
	}
	return dummy.Next
}

// 和 v2 版本比较，这个优点在哪里？
func removeElementsV3(head *ListNode, val int) *ListNode {
	ret := &ListNode{Val: 0, Next: head}
	index := ret
	for index.Next != nil {
		if index.Next.Val == val {
			next := index.Next
			index.Next = next.Next
		} else {
			index = index.Next
		}
	}
	return ret.Next
}