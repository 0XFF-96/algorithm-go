package linklist

// 1290. Convert Binary Number in a Linked List to Integer
func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}

	// 从 101 和 100 两种情况讨论
	// 1>>2*1 + 1>>1*0 + 1>>0*1
	// 1>>2*1 + 1>>1*0 + 1>>0*0
	// 要知道这些信息才能够啊。
	// value := getDecimalValue(head.Next)
	// return 1 >> value

	sum, _ := helper(head)
	return sum
}

func helper(head *ListNode) (int, int) {
	if head == nil {
		return 0, 0
	}

	sum, level := helper(head.Next)
	return sum + (1 << level) * head.Val, level +1
}