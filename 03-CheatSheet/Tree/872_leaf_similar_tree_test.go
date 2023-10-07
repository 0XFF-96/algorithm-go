package tree

import (
	"math"
	"testing"
	"time"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	var helper func(root *TreeNode, fn func (cur *TreeNode))

	helper = func(root *TreeNode, fn func(cur *TreeNode)) {
		if root == nil {
			return
		}
		helper(root.Left, fn)
		helper(root.Right, fn)
		if root.Left == nil && root.Right == nil {
			fn(root)
		}
		return
	}

	var leafSum1 []int
	var leafSum2 []int
	helper(root1, func(cur *TreeNode){
		leafSum1 = append(leafSum1, cur.Val)
	})
	helper(root2, func(cur *TreeNode){
		leafSum2 = append(leafSum2, cur.Val)
	})

	if len(leafSum1) != len(leafSum2) {
		return false
	}
	for i:=0; i < len(leafSum1); i++{
		if leafSum1[i] != leafSum2[i]{
			return false
		}
	}
	return true
}

func TestTime(t *testing.T){
	//
	//t.Log(int(time.Hour * 24 * 30 / time.Second))
	//t.Log(36002430/86400)
	//t.Log(int(time.Hour * 24 * 31 / time.Second))


	t.Log(int(time.Hour * 24 * 1 / time.Second))  // 86400
	t.Log(int(time.Hour * 24 * 3 / time.Second))  // 259200
	t.Log(int(time.Hour * 24 * 7 / time.Second))  // 604800
	t.Log(int(time.Hour * 24 * 30 / time.Second)) // 259200
	t.Log(int(time.Hour * 24 * 365 / time.Second)) // 31536000

	t.Log(math.MaxInt32)
}


