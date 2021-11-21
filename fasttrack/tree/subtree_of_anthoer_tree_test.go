package tree



// 标准答案
// 能否用递归的解法解出来？
// 基本条件: same tree 的判断：
// same tree 递归和非递归解法： https://www.youtube.com/watch?v=ySDDslG8wws
//


// 高级的办法
// https://leetcode.com/problems/subtree-of-another-tree/discuss/102741/Python-Straightforward-with-Explanation-(O(ST)-and-O(S%2BT)-approaches)


//func isSubtree(s *TreeNode, t *TreeNode) bool {
//	if s == nil && t == nil {
//		return true
//	}
//	if s != nil && t == nil {
//		return false
//	}
//	if s.Val == t.Val {
//		return isSubTree(s.Left, t.Left) && isSubTree(s.Right, t.Right)
//	}
//
//	return isSubTree(s.Left, t) || isSubTree(s.Right, t)
//}


// 参考一下这里的答案
// https://leetcode.com/problems/subtree-of-another-tree/discuss/407630/Golang-Recursive-solution
//



// ugly solution
// https://leetcode.com/problems/subtree-of-another-tree/discuss/102788/Golang-solution-serialization-of-two-trees


// =======
// 下面的解法，不行...
//
//func isSubtreeV2(s *TreeNode, t *TreeNode) bool {
//	if s == nil && t == nil {
//		return true
//	}
//	if s == nil || t == nil {
//		return false
//	}
//	if isSameTree(s, t) {
//		return true
//	}
//
//	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
//}



func isSameTreeV22(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	// 这一句是必须的
	// 否则会导致 nil pointer reference
	if s == nil || t == nil {
		return false
	}

	// 有一行代码
	// 没有改正过来...
	//
	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right) && s.Val == t.Val
}
// ===============


// merkel 方法
// https://leetcode.com/problems/subtree-of-another-tree/discuss/102741/Python-Straightforward-with-Explanation-(O(ST)-and-O(S%2BT)-approaches)


// 递归解法2
// https://leetcode.com/problems/subtree-of-another-tree/discuss/443079/Golang-15-lines-of-solution-with-explanations


// https://www.youtube.com/watch?v=73PQ9raLEVs&t=3s
// youtube 视频解法
// 有什么技巧？
// 难点？

