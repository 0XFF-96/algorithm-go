package tree

import (
	"strconv"
)

// https://www.youtube.com/watch?v=73PQ9raLEVs&t=3s
// 暂时不会
// 有什么思维方法可以解决
// 提示点：序列化、map、相同结构


// 快速查找一个树的结构是否在里面。 map
//
// 视频1：
// https://www.youtube.com/watch?v=nxHBg7hm0rs [已看]
// https://www.youtube.com/watch?v=JLK92dbTt8k [已看]
// https://www.youtube.com/watch?v=LYU3y0-59_k [视频]

// [0,0,0,0,null,null,0,null,null,null,0]
// 上面这种情况，没有考虑过
//
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := map[string]int{}

	var ret []*TreeNode

	var helper func(root *TreeNode) string
	helper = func(root *TreeNode) string {
		if root == nil { return "#"}

		// serial 序列化的顺序对与这道题，会不会有影响？
		// 会？
		// 为什么？！！！
		// 如果是 helper(root.Left) + strconv.Itoad + helper(root.Right) 就不行了
		// 根左右
		serial :=   strconv.Itoa(root.Val)+ helper(root.Left)  + helper(root.Right)

		_, ok := m[serial]
		if ok {
			m[serial]++
		} else {
				m[serial] = 1
		}
		// 恰好两个才 add, 否则会重复
		if m[serial] == 2 {
			ret = append(ret, root)
		}
		return serial
	}

	helper(root)
	return ret
}

//Time Complexity: O(N^2)O(N
//2
//), where NN is the number of nodes in the tree. We visit each node once, but each creation of serial may take O(N)O(N) work.
//
//Space Complexity: O(N^2)O(N
//2
//), the size of count.

// 后序遍历的解法
// https://leetcode.com/problems/find-duplicate-subtrees/discuss/187242/Golang-postorder-%2B-hashmap-without-global-variables.
// 还有没有其他的解法？
// 如果分析时间和空间复杂度？