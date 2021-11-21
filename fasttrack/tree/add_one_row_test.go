package tree

import (
	"testing"
)

func TestAddRow(t *testing.T){
	// 正确解法
}

// 答案
// https://leetcode.com/articles/add-one-row-in-a-tree/
// 暂时不知道

// 移动要熟练这几种解法
// 1、
// 2、
// 3、
// 4、

// 层次遍历的思想
// 从 d-1 层出发

// 为什么是错误的呢？
func addOneRowWrong(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		root.Left = &TreeNode{Val:d}
		root.Right = &TreeNode{Val:d}
		return root
	}

	queue := []*TreeNode{root}

	height := 0
	for len(queue) != 0 {
		height++
		size := len(queue)
		for i:=0; i < size; i++{
			top := queue[0]
			queue = queue[1:]
			if height == d-1 {
				tmpLeft := top.Left
				tmpRight := top.Right
				top.Left = &TreeNode{Val:v}
				top.Right = &TreeNode{Val:v}
				top.Left.Left = tmpLeft
				top.Right.Right = tmpRight
			}
		}
	}

	return root
}


// 层次遍历的思想
// 从 d-1 层出发
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		// tmpLeft := root.Left
		// tmpRight := root.Right
		// root.Left = &TreeNode{Val:d}
		// root.Right = &TreeNode{Val:d}
		// root.Left.Left = tmpLeft
		// root.Right.Right = tmpRight
		// return root

		tmp := &TreeNode{Val:v}
		tmp.Left = root
		return tmp
	}

	queue := []*TreeNode{root}
	height := 0
	for len(queue) != 0 {
		height++
		size := len(queue)
		for i:=0; i < size; i++{
			top := queue[0]
			queue = queue[1:]

			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			if height == d-1 {
				tmpLeft := top.Left
				tmpRight := top.Right
				top.Left = &TreeNode{Val:v}
				top.Right = &TreeNode{Val:v}
				top.Left.Left = tmpLeft
				top.Right.Right = tmpRight
			}
		}
	}

	return root
}