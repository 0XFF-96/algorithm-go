package tree

import (
	"testing"
)

func TestMaxDepth(t *testing.T){
	root := &TreeNode{Val:1}
	root.Right = &TreeNode{Val:2}
	root.Right.Right = &TreeNode{Val:3}

	m := maxDepth(root)

	m2 := maxDepthDFS(root)
	t.Log(m, m2)
}

// 代码更清晰的解法，https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/solution/tu-jie-104-er-cha-shu-de-zui-da-shen-du-di-gui-pyt/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right  := maxDepth(root.Right)

	if left > right {
		return left +1
	} else {
		return right +1
	}
}

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/solution/golangdie-dai-qiu-er-cha-shu-de-zui-da-shen-du-by-/
// 启发式算法
func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0

	for len(queue) != 0 {
		size := len(queue)
		for i:=0; i<size;i++{
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			queue = queue[1:]
		}
		depth++
	}
	return depth
}
