package tree


// Maximum Width of Binary Tree
// https://www.youtube.com/watch?v=NA8B84DZYSA&t=26s

// 不懂怎么计算 width
// height
// https://www.youtube.com/watch?v=tRH8RS99wPk
// 找到节点于节点的对应数字关系
//

// 总结：
// 这是一题找规律，利用树节点，之间信息计算的题目.
// 归类于树的性质：树高、树宽、树的最大深度、树的最短长度、pathSum 等题目
// position value

// The main idea in this question is to give each node a position value.
	// If we go down the left neighbor, then position -> position * 2;
// and if we go down the right neighbor, then position -> position * 2 + 1.
// This makes it so that when we look at the position values L and R of two nodes with the same depth,
// the width will be R - L + 1.


// how to use this info to tackle the problem ?
// The binary tree has the same structure as a full binary tree
// vertical projection of a binary tree

// 前序遍历的框架
func widthOfBinaryTree(root *TreeNode) int {
	ans := 0

	// depth:leftmost node's position
	leftmostMap := map[int]int{}
	var helper func(root *TreeNode, depth int, pos int)

	helper = func(root *TreeNode, depth int, pos int){
		if root == nil {
			return
		}

		// only set once
		if _, ok := leftmostMap[depth]; !ok {
			leftmostMap[depth] = pos
		}

		ans = max(ans, pos-leftmostMap[depth] +1)
		helper(root.Left, depth+1, 2 *pos)
		helper(root.Right, depth+1, 2 *pos+1)
	}

	helper(root, 0, 0)
	return ans
}

// https://leetcode.com/problems/maximum-width-of-binary-tree/
// how to use this info to tackle the problem ?
// The binary tree has the same structure as a full binary tree
// vertical projection of a binary tree

// 难点
// 1、如何计算宽度 2、这题目需要的格外信息有哪些？ 3、有两个解题方法
// 给出任意例子，能够计算树的宽度。
// 栈遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 终于都 AC 了
// 爽😊
//func widthOfBinaryTreeV2(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	queue := []*Anode{&Anode{node:root, pos:0}}
//	curDepth := 0
//	ans := 0
//	left := 0
//	for len(queue) != 0 {
//		size := len(queue)
//		for i:=0;i<size;i++{
//			top := queue[0]
//			queue = queue[1:]
//
//			if top.node != nil {
//				queue = append(queue, &Anode{node:top.node.Left, depth: top.depth +1, pos: top.pos*2})
//				queue = append(queue, &Anode{node:top.node.Right, depth: top.depth +1, pos: top.pos*2 +1})
//
//				if curDepth != top.depth {
//					curDepth = top.depth
//					left = top.pos
//				}
//
//				ans = max(top.pos - left + 1, ans)
//			}
//
//		}
//
//	}
//	return ans
//
//}

//type Anode struct {
//	node *TreeNode
//	pos int
//	depth int
//}
//
//
//func max(a, b int)int{
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
