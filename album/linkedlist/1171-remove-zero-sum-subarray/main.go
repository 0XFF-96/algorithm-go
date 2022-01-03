package _171_remove_zero_sum_subarray

// 1、https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/discuss/388462/golang-94.5100-using-hashmap-easy-to-understand
// 2、情况:
// 3、

// 文章2： https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/discuss/388462/golang-94.5100-using-hashmap-easy-to-understand

// brute force 的算法
// 基本上算一个比较完美的算法了
// 有一个重要的结论，prefix sum 出现相等时，证明其中的数组元素有相同的 subsequence 和
// 用 set 来记录有没有出现过相同的 prefix sum
func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	finalHead := &ListNode {
		Val: 0,
		Next: head,
	}
	currentHead := finalHead // remove nodes from currentHead
	for ; currentHead != nil; {
		currentSum := 0
		node := currentHead.Next
		for ; node != nil; {
			currentSum += node.Val
			if currentSum == 0 {
				currentHead.Next = node.Next
			}
			node = node.Next
		}
		currentHead = currentHead.Next
	}
	return finalHead.Next
}