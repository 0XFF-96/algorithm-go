package tree

import (
	"strconv"
)

// Solution
// 这一题和 find duplicates tree 有异曲同工之妙
// https://leetcode.com/problems/construct-string-from-binary-tree/solution/
// 有四种情况需要考虑

// 序列化树的相关操作。
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if (t.Left == nil && t.Right == nil ) {
		return strconv.Itoa(t.Val) + ""
	}
	if t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	}
	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) +")" + "(" +tree2str(t.Right) + ")"
}

// 递归版本
// 亲自写出来
// https://leetcode.com/problems/construct-string-from-binary-tree/discuss/265083/golang-iterative-dfs
// https://www.geeksforgeeks.org/find-duplicate-subtrees/
// http://ayushcshah.github.io/algorithm/binarytree/2016/04/01/detect-duplicate-subtrees.html