package archive_Tree

import (
	"testing"

	"pratice-go/algorithmTopic/tree"
)

// https://leetcode.com/problems/most-frequent-subtree-sum/
// 这题，连题目都读不懂...
// 这道题目的解决方法
// https://www.youtube.com/watch?v=HWQnh50t1mw
//
// 三个变量：map, 全局最大 frequent, 后序遍历的思想
// 左右子树都遍历了才能行
//

// 后序遍历的非递归解法？
// 模板？

// 答案伪代码描述（两个点）
// 1、用 map 记录
// 2、后序遍历

// subtree sum
// case1:
// case2:
// case3:
// 有没有非递归的解法？
// 时间和空间复杂度是多少？
func TestFindFrequent(t *testing.T){
	tree1 := tree.initTree()

	res := findFrequentTreeSum(tree1)
	t.Log(res)
}


func findFrequentTreeSum(root *tree.TreeNode) []int {
	m := map[int]int{}
	max := 0

	findFreqHelper(root, m, &max)

	var res []int
	for k, v := range m {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func findFreqHelper(root *tree.TreeNode, m map[int]int, max *int)int{
	if root == nil {
		return 0
	}

	left := findFreqHelper(root.Left, m, max)
	right := findFreqHelper(root.Right, m, max)

	sum := left + right + root.Val
	m[sum]++
	if m[sum] > *max {
		*max = m[sum]
	}
	return sum
}