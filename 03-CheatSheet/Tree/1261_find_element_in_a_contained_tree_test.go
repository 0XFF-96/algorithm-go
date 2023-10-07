package tree

import (
	"testing"
)

type FindElements struct {
	Val int
	Left, Right *FindElements
}


// 速度太慢了
// 从哪里可以优化一下？
// Runtime: 944 ms, faster than 6.12% of Go online submissions for Find Elements in a Contaminated Binary Tree.
// Memory Usage: 7.7 MB, less than 100.00% of Go online submissions for Find Elements in a Contaminated Binary Tree.
//

// 优化点1： 1、查找算法-》 O(1) 复杂度。使用 map 来进行优化，https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/discuss/556211/Go-solution-(dfs)
// 优化点2： bit path https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/discuss/431362/Java-bit-path-Time-%3A-O(logn)
// 优化点3： binaryPath without set:https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/discuss/431283/Python-binary-path-without-set
//

func TestConstructor(t *testing.T){
	root := &TreeNode{Val:-1}
	root.Right = &TreeNode{Val:-1}
	//root.Left = &TreeNode{Val:-1}
	//root.Left.Left = &TreeNode{Val:-1}
	//root.Left.Right = &TreeNode{Val:-1}
	rootF := ConstructorVV1(root)
	t.Log(rootF.Left)
	t.Log(rootF.Right.Val)
	// t.Log(rootF.Left.Left)
	t.Log(rootF.Find(2))
}

func ConstructorVV1(root *TreeNode) FindElements {
	// rootF := FindElements{Val:0}
	// 遍历树的算法
	var helper func(root *TreeNode, val int) *FindElements

	helper = func(root *TreeNode,val int)*FindElements {
		if root == nil {
			return nil
		}
		rootF := &FindElements{Val: val}
		rootF.Left = helper(root.Left, val*2 +1)
		rootF.Right = helper(root.Right, val*2 +2)

		return rootF
	}

	rootF := helper(root, 0)

	return *rootF
}

// 二次叉搜索
func (this *FindElements) Find(target int) bool {
	if this == nil {
		return false
	}
	cur := this

	// 常量放在 等于号哪一边？
	// 这个是 binary tree, 不是 bst 不能用这个表示
	//for cur != nil {
	//	fmt.Println(cur.Val)
	//	if cur.Val == target {
	//		return true
	//	}
	//	if cur.Val > target {
	//		cur = cur.Right
	//	} else {
	//		cur = cur.Left
	//	}
	//}
	return find(cur, target)
}

// 需要用这个重构一下
//
func find(root *FindElements, target int) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	return find(root.Left, target) || find(root.Right, target)
}