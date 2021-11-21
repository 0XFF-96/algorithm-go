package tree

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/solution/
// 有更多广度和深度题目的算法.
//
func TestFindTarget(t *testing.T){
	// tree1 := initTree()
	nodes := &TreeNode{Val:2}
	nodes.Left = &TreeNode{Val:1}
	nodes.Right = &TreeNode{Val:3}

	r := findTarget(nodes, 3)
	t.Log(r)
}


// ac
func findTarget(root *TreeNode, k int) bool {
	m := map[int]struct{}{}

	var helper  func(root *TreeNode) bool

	helper = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		tmp := k - root.Val
		fmt.Println(tmp)
		_, ok := m[tmp]
		if ok {
			return true
		}
		m[root.Val] = struct{}{}

		return helper(root.Left) || helper(root.Right)
	}

	return helper(root)
}
