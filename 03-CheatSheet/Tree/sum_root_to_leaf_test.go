package tree

import (
	"fmt"
	"testing"
)

func sunNumbers(root *TreeNode) int {

	// findLeafSum()[][]int
	// how to record the path here
	// hasPathSum
	// add them up
	// return the result.
	return 0
}

func TestSumNumbers(t *testing.T){
	// tree1 := initTree()

	//sum := rootToLeafPaths(tree1)
	//t.Log(sum)

	tree2 := &TreeNode{Val:1}
	tree2.Right = &TreeNode{Val:2}
	tree2.Left = &TreeNode{Val:3}
	t.Log(sumNumbers(tree2))
}


// 递归解法
// 错误的...
func rootToLeafPaths(root *TreeNode)[][]int{
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	right := rootToLeafPaths(root.Left)
	left  := rootToLeafPaths(root.Right)

	// merge right and left
	res := append(right, left...)

	// res
	// add cur Node value
	// 这里要进行一次深拷贝
	var ret [][]int
	for _, r := range res {
		r = append(r, root.Val)
		tmp := []int{}
		fmt.Printf("%v", r)
		copy(tmp, r)
		fmt.Printf("%v", tmp)
		ret = append(ret, tmp)
	}
	return ret
}

// 迭代解法1
//


// 这题类似于 PathSum
// 参考相关实现

// 中序遍历有 bug 的原因
// 明白了
// 在抵达右叶子节点前，根就在栈里被弹出来了
func sumNumbersW(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var pathSum []int

	stack := []*TreeNode{}
	cur := root

	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			// 开始了...
			//if cur.Left == nil && cur.Right == nil {
			//	sum := 0
			//	for i := len(stack)-1;i >=0; i-- {
			//		sum = sum * 10
			//		sum += stack[i].Val
			//	}
			//	pathSum = append(pathSum, sum)
			//}

			stack = stack[:len(stack)-1]
			cur = cur.Right
			continue
		}
		stack = append(stack, cur)
		if cur.Left == nil && cur.Right == nil {
			sum := 0
			for i := len(stack)-1;i >=0; i-- {
				sum = sum * 10
				sum += stack[i].Val
			}
			pathSum = append(pathSum, sum)
		}
		cur = cur.Left
	}

	fmt.Printf("%v", pathSum)
	ret := 0
	for _, v := range pathSum {
		ret += v
	}
	return ret
}


func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var pathSum []int
	var helper = func(root*TreeNode,path []int){}
	helper = func(root *TreeNode, path []int){
		if root == nil {
			return
		}
		path = append(path, root.Val)
		helper(root.Left, path)
		helper(root.Right, path)
		if root.Left == nil && root.Right == nil {
			sum := 0
			for _, v := range path {
				sum *= 10
				sum += v
			}
			pathSum = append(pathSum, sum)
		}
		return
	}

	helper(root, []int{})
	ret := 0
	for _, v := range pathSum {
		ret += v
	}
	return ret
}


// 高级一点
// 看看利用栈怎么解题？
// https://leetcode.com/problems/sum-root-to-leaf-numbers/discuss/330633/Golang-DFS-%2B-Stack-solution