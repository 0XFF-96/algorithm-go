package tree

import (
	"testing"
)

func TestLeverOrder(t *testing.T){
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:3}

	res := levelOrderDFS(root)

	res2 := levelOrderWrong(root)
	t.Log(res,res2)

	res4 := levelOrderWithQueue(root)
	t.Log(res4)

}

// 两栈轮流倒...
// 空间复杂度 2n
// 时间复杂度 n
// 有一点非常不友好的是，算法题的命名不知道怎么写
// 导致有时间看错参数...
func levelOrderWrong(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	stack2 := []*TreeNode{}
	var ret [][]int
	// 广度优先算法的框架

	// 这个for 的时间消耗等于树的高度
	for len(stack) != 0 || len(stack2) != 0 {
		// 两个栈
		if len(stack) != 0 {
			var res []int
			for _, s := range stack {
				res = append(res, s.Val)
				if s.Left != nil {
					stack2 = append(stack2, s.Left)
				}
				if s.Right != nil {
					stack2 = append(stack2, s.Right)
				}
			}
			stack = nil
			ret = append(ret, res)
		}

		if len(stack2) != 0 {
			var res []int
			for _, s := range stack2 {
				res = append(res, s.Val)
				if s.Left != nil {
					stack = append(stack, s.Left)
				}
				if s.Right != nil {
					stack = append(stack, s.Right)
				}
			}
			stack2 = nil
			ret = append(ret, res)
		}
	}
	return ret
}


func levelOrderWithQueue(root *TreeNode)[][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	queue := []*TreeNode{root}
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
		ret = append(ret, level)
	}
	return ret
}


// 先写一个递归的算法
// 记录层数
// DFS Solution
func levelOrderDFS(root *TreeNode)[][]int{
	if root == nil {
		return nil
	}
	result := [][]int{}
	levelOrderHelper(root, 0, &result)
	return result
}

func levelOrderHelper(root *TreeNode, level int, result *[][]int){
	if root == nil {return }
	if level == len(*result) {*result = append(*result, []int{})}

	(*result)[level] = append((*result)[level], root.Val)

	levelOrderHelper(root.Left, level+1, result)
	levelOrderHelper(root.Right, level+1, result)
}


// TODO: level order with iteration
// https://leetcode.com/problems/binary-tree-level-order-traversal/discuss/33541/Golang-solution-(9ms)
// 非递归版本的 DFS
// level order 的另一种做法
// https://leetcode.com/problems/binary-tree-level-order-traversal/discuss/33541/Golang-solution-(9ms)


// Binary Tree Level Order
func levelOrderBottom(root *TreeNode) [][]int{
	if root == nil {
		return nil
	}
	// rec := map[int][]int{}
	// 这种初始化方法不可靠...
	var rec = make(map[int][]int)
	level := 0
	levelOrderByMap(root, rec, level)
	var result [][]int
	for i := len(rec)-1; i >= 0; i--{
		result = append(result, rec[i])
	}

	return nil
}

func levelOrderByMap(root *TreeNode, rec map[int][]int, level int){
	if root == nil {
		return
	}
	rec[level] = append(rec[level], root.Val)

	level++
	levelOrderByMap(root.Left, rec, level)
	levelOrderByMap(root.Right, rec, level)
}