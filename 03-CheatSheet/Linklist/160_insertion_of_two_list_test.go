package linklist


// 有三种解法
//
//

// Here's my solution which is like approach 3 but a little different.
// I store the size of ListA and ListB as len1 and len2.
// Then I reset the pointers to headA and headB
// and find the difference between len1 and len2,
// and then let the pointer of the longer list proceed
// by the difference between len1 and len2.
// Finally, traverse through the lists again,
// the intersection node can be easily found

//
// Runtime: 40 ms, faster than 83.44% of Go online submissions for Intersection of Two Linked Lists.
// Memory Usage: 8.4 MB, less than 20.00% of Go online submissions for Intersection of Two Linked Lists.


//follow-up 的内容
// - If the two linked lists have no intersection at all, return null.
// - The linked lists must retain their original structure after the function returns.
// - You may assume there are no cycles anywhere in the entire linked structure.
// - Your code should preferably run in O(n) time and use only O(1) memory.

// 自己 AC 的写法...
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	countA := lenList(headA)
	countB := lenList(headB)

	curA := headA
	curB := headB

	if countA > countB {
		n := countA - countB
		for n > 0 {
			curA = curA.Next
			n--
		}
	} else {
		n := countB - countA
		for n > 0 {
			curB = curB.Next
			n--
		}
	}

	for curA != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}

func lenList(l *ListNode) int{
	count := 0
	for l != nil {
		count++
		l = l.Next
	}
	return count
}


