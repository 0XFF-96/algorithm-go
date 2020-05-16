package _019_next_greater_value


// We can use a stack that stores nodes in monotone
// decreasing order of value.
// When we see a node_j with a larger value, * ？
// every node_i in the stack has next_larger(node_i) = node_j * ？ 怎么理解？
// 简单版本
// n平方， O(N) 的算法
func nextLargerNodes(head *ListNode) []int {
	var res []int
	for curr := head; curr != nil; curr = curr.Next {
		res = append(res, nextLargerNode(curr))
	}
	return res
}

func nextLargerNode(head *ListNode) int {
	for p := head; p != nil; p = p.Next {
		if p.Val > head.Val {
			return p.Val
		}
	}
	return 0
}


// 用可视化的方法考虑问题
// [2, 7, 3, 4, 5]
// i | A[i] | stack | ans
// monotonic stack
// 比我小的元素不可能构成一个解法
//

// 从上一个版本中迭代
// 如何能够快速查找 nextLargerNode
// 对于这个点进行相关优化才行
// 根据空间换时间的定理，一定可以利用数据结构优化查找的过程
//