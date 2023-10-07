package tree

// 层次遍历的一种解法
// 相关条件


// 1、
// 2、

// 层次遍历的解法 + 找到每层的最大值
//

import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var ret []int

	for len(queue) != 0 {
		size := len(queue)-1
		maxLevelVal := math.MinInt32

		for i:=0;i<=size;i++{
			top := queue[0]
			queue = queue[1:]
			if top.Val > maxLevelVal {
				maxLevelVal = top.Val
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			if i == size {
				ret = append(ret, maxLevelVal)
			}
		}

	}
	return ret
}

//