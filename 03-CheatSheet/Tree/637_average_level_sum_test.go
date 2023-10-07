package tree

import (
	"fmt"
	"testing"
)


// 层次遍历的解法
// BFS 和 DFS 两种都要掌握
// https://leetcode.com/problems/average-of-levels-in-binary-tree/solution/
func TestAverageOfLevels(t *testing.T){
	tree1 := initTree()
	sum := averageOfLevels(tree1)
	t.Log(sum)
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var ret []float64
	for len(queue) != 0 {
		size := len(queue)
		levelSum := 0.0
		for i := 0;i<size;i++{
			top := queue[0]
			queue = queue[1:]

			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}

			levelSum += float64(top.Val)
		}
		levelAverage := levelSum / float64(size)
		fmt.Println(levelAverage)
		ret = append(ret, float64(levelAverage))
	}

	return ret
}