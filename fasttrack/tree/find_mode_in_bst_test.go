package tree

import (
	"math"
	"testing"
)

// Find Mode in Binary Search Tree
// https://www.youtube.com/watch?v=1FJDyBSfEbo
// 题目相关： https://leetcode.com/problems/find-mode-in-binary-search-tree/


// 需要 4 个不同的参数
// 记录遍历状态下的不同反应
// curVal, curCount, maxCount, modeCount
// 一个数组

// 虽然写错了，
// 但是有几个方面写好了
// 1、Morii 的遍历方式
// 2、函数式编程
// 3、记录运行时需要的参数


func TestFindMode(t *testing.T){
	tree := initTree()

	// 这里少了最后的边界节点 7
	// 是在哪里漏判断 7 ？
	sum := findModeWrong(tree)
	t.Log(sum)


	//node := &TreeNode{Val:1}

	// 有没有办法从 第 2 个节点开始判断？
	//
	mode := findMode(tree)
	t.Log(mode)

}

func findModeWrong(root *TreeNode) []int{
	if root == nil {
		return nil
	}

	var curCount, curVal, maxCount, modeCount int
	var modes []int
	var inorder = func(root *TreeNode, fn func(cur *TreeNode)){}

	inorder = func(root *TreeNode,  fn func(cur *TreeNode)) {
		if root == nil {
			return
		}
		var predessor *TreeNode
		cur := root

		for cur != nil {
			if cur.Left == nil {
				// do something
				// how to extract this method to it
				// inject the method to it.
				fn(cur)
				cur = cur.Right
			} else {
				predessor = cur.Left
				for predessor.Right != cur && predessor.Right != nil {
					predessor = predessor.Right
				}

				if predessor.Right == nil {
					// hav't been visited yet
					predessor.Right = cur
					cur = cur.Left
				} else {
					predessor.Right = nil
					fn(cur)
					cur = cur.Right
				}
			}
		}
	}

	// run
	inorder(root, func(cur *TreeNode){

		// 这些判断语句有错误
		if curVal != cur.Val {
			if curCount > maxCount {
				maxCount = curCount
				modes = []int{curVal}
				modeCount++
			} else if curCount == maxCount {
			    modeCount++
				modes = append(modes, curVal)
			}

			curCount = 1
			curVal = cur.Val
		} else {
			// equal
			// 要不要判读mod
			curCount ++
		}
	})
	return modes
}

// https://leetcode.com/problems/find-mode-in-binary-search-tree/discuss/279344/Golang-using-map

// https://leetcode.com/problems/find-mode-in-binary-search-tree/discuss/279344/Golang-using-map
//

// 看看这两个方法有什么优缺点...
// https://leetcode.com/problems/find-mode-in-binary-search-tree/discuss/407458/Go-golang-two-solutions


// 用 map 实现
// 1、用 [modeCount] modes
// 2、在中序遍历的时候记录几个信息

func findMode(root *TreeNode) []int{
	modes := map[int][]int{}
	curCount := 0
	maxModeCount := 0
	preVal := math.MinInt64

	var helper = func(root *TreeNode, fn func(root *TreeNode)){}

	helper = func(root *TreeNode, fn func(root *TreeNode)){
		if root == nil {
			return
		}
		helper(root.Left, fn)
		fn(root)
		helper(root.Right,fn)
		return
	}

	helper(root, func(root *TreeNode){
		if root == nil {
			return
		}

		// 这里的判断顺序不对
		// 导致后果...
		if root.Val != preVal {
			// 更新信息
			if curCount > maxModeCount {
				maxModeCount = curCount
				modes[curCount] = []int{preVal}
			} else if curCount == maxModeCount {
				modes[curCount] = append(modes[curCount], preVal)
			}

			// 清零数据
			preVal = root.Val
			curCount = 0
		} else {
			curCount ++
		}
	})

	return modes[maxModeCount]
}


// 难点1：
// 为什么 count 初始化为 0 而 max 初始化为 1 ？
// 通过制造差异，从而避开第一个节点的时刻。
// 难点2：
// *max = *count // Set max
// 忘记设置这个
func findModeGood(root *TreeNode)[]int{
	if root==nil {return nil}

	pre := -1<<31
	count := 0
	max := 1

	// 这个是一个数组.
	result := make([]int,0)
	findModeHelper(root, &pre, &count, &max, &result)

	return result
}

func findModeHelper(root *TreeNode, pre *int, count *int, max *int, result *[]int){
	if root == nil {
		return
	}
	findModeHelper(root.Left, pre, count, max, result)

	if root.Val == *pre {*count++} else {*count = 1}

	if *max < *count {
		*result = []int{root.Val}
	} else if *max == *count {
		*result = append(*result, root.Val) // Append to result
	}
	*pre = root.Val

	findModeHelper(root.Left, pre, count, max, result)
}
