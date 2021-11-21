package tree

import (
	"testing"
)

func TestZigZagLevel(t *testing.T){
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:9}
	root.Right = &TreeNode{Val:20}
	root.Right.Right = &TreeNode{Val:7}
	root.Right.Left = &TreeNode{Val:15}

	z := zigzagLevelOrder(root)
	t.Log(z)
}

// 答案和这个很类似
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/discuss/33906/Golang-iterative-2-solutions-using-reverse-and-two-stacks
func zigzagLevelOrder(root *TreeNode)[][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	queue := []*TreeNode{root}

	depth := 0
	for len(queue) != 0 {
		size := len(queue)
		var level []int

		// level travel
		for i:=0;i < size;i++{

			// 取出队列头部
			cur := queue[0]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			level = append(level, cur.Val)
			queue = queue[1:]
		}

		// 在这里增加 depth 信息
		// 然后将 level 的数组反序
		// 代码会更加清晰

		if depth % 2 != 0 {
			reverse(level)
		}
		ret = append(ret, level)
		depth ++
	}
	return ret
}

func reverse(data []int)  {
	l:= len(data)
	for i:=0;i< l/2;i++{
		data[i],data[l-1-i] =data[l-1-i],data[i]
	}
}