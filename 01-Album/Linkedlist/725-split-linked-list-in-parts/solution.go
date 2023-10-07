package main

type ListNode struct {
	Val int
	Next *ListNode
}

// If there are N nodes in the list, and k parts, then every part has // N/k elements, except the first N%k parts have an extra one
// the thing is we don't know what N is until we count it
//
// 有相关参考答案: https://leetcode.com/problems/split-linked-list-in-parts/
func splitListToParts(root *ListNode, k int) []*ListNode {
	// 类似的 split 题目
	// 快慢指针？
	//

	cur := root

	N := 0
	for cur != nil {
		cur = cur.Next
		N++
	}

	width := N / k
	rem := N % k

	ret := make([]*ListNode, k)
	cur = root
	var carry int
	for i:=0; i < k;i++{
		head := &ListNode{}
		write := head

		// width + ( i < rem ? 1:0)
		// 这里的含义？
		// 相差不超过 1
		//
		if i < rem {
			carry = 1
		} else {
			carry = 0
		}

		for j:=0;j < width + carry; j ++{
			write.Next = &ListNode{Val: cur.Val}
			write = write.Next

			// 为什么要放在后面？
			// 不会超
			// 因为有 k 保护
			// if cur != nil {
			//     cur = cur.Next
			// }
			// 是的，因为 k 和 width 都是
			// 根据链表的长度计算好的数值
			// 所以，没有 cur != nil
			// 这个保护也没什么太大问题。
			cur = cur.Next
		}

		ret[i] = head.Next
	}
	return ret
}
