package tree

import (
	"math"
	"testing"
)


// 其实和 530 是同一题目

// 怎么利用递归解题？
// 记录前一个元素
func TestInt64(t *testing.T){

	// 溢出了...
	t.Log(math.MinInt64)
}

func minDiffInBST(root *TreeNode) int {
	min := math.MaxInt64
	cur := root
	stack := []*TreeNode{}


	// 又是遇到初始化的难题
	//
	pre := -10000000

	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 访问
			tmp := cur.Val - pre
			if tmp < min {
				min = tmp
			}
			pre = cur.Val
			cur = cur.Right
			continue
		}

		stack = append(stack, cur)
		cur = cur.Left
	}
	return min
}
