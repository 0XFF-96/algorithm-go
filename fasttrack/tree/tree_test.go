package tree

import (
	"testing"
)

// https://leetcode.com/tag/tree/
// 1、在 5 天内过一下这里面的 tag

// 总结
// 学到了一种新的方法
// https://stackoverflow.com/questions/5502916/explain-morris-inorder-tree-traversal-without-using-stacks-or-recursion

// TODO: binary Tree Inorder Traversal
// https://leetcode.com/problems/binary-tree-inorder-traversal/


// TODO: https://leetcode.com/problems/path-sum-ii/
// https://leetcode.com/problems/path-sum-ii/discuss/36840/20ms-Golang-solution
// 1、利用闭包函数的优势
// 2、深拷贝
// 用思维框架的方式去思考算法题


// 不会 maximum path
// https://leetcode.com/problems/binary-tree-maximum-path-sum/



//  lowest common ancenstor 是最合适的解法
//  尝试从各个方面去进行求解？
//
// TODO: 后序遍历的非递归解法...


// TODO: Symmetric Tree
// 解法的伪代码？
// Symmetric Tree 的解法有点卡
// https://leetcode.com/problems/symmetric-tree/
// 大概知道复制一份节点信息



// TODO: Sum Root To Leaf Numbers
// https://leetcode.com/problems/sum-root-to-leaf-numbers/
//


// TODO: house_robber 结合动态规划的题目，暂时停止
// https://leetcode.com/problems/house-robber/submissions/


// TODO: diameterOfBinaryTree
// sum up tree 中的部分路径
// 这部分的题目还不怎么熟悉需要加强练习一下


// TODO: construct string from a bianry tree
// Construct String from Binary Tree


// TODO: 熟练 Add row 的几种解法
// https://leetcode.com/articles/add-one-row-in-a-tree/
//


// TODO: 998 maximum_bianry_tree
// 这个一定要有，特别是几种不同的解法。
// 逻辑的合并
//
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func inorderTraversal(root *TreeNode)[]int{
	var res []int

	if root == nil {
		return res
	}

	if root.Left != nil {
		res = inorderTraversal(root.Left)
	}
	res = append(res, root.Val)
	right := inorderTraversal(root.Right)

	for _, e := range right {
		res = append(res, e)
	}
	return res
}

func inorderTraversalWithStack(root *TreeNode)[]int{
	var result []int

	stack := []*TreeNode{}
	cur := root

	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 已经访问当前节点了
			result = append(result, cur.Val)

			// 为什么？下面两行是必须的？
			cur = cur.Right
			continue
		}
	stack = append(stack, cur)
	cur = cur.Left
	}
	return result
}

//func preorderTraversal(root *TreeNode)[]int{
//	var res []int
//	if root == nil {
//		return res
//	}
//	res = append(res, root.Val)
//	res = append(res, preorderTraversal(root.Left)...)
//	res = append(res, preorderTraversal(root.Right)...)
//	return res
//}

// 在哪里构建了一条回溯的边？
// 什么时候进行分支转向？

//// 脑海中要想像有一个图形
//// https://www.youtube.com/watch?v=wGXB9OWhPTg
//func MorriInorderTravel(root *TreeNode)[]int{
//	var res []int
//	current := root
//	for current != nil {
//	if current.Left == nil {
//		res = append(res, current.Val)
//		current = current.Right
//	} else {
//		// find the predecessor.
//		// To Find the predecess keep going right
//		// till right node is not null or right node
//		// is not current.
//
//		// 这里找 predecessor 的前提是，左分支不为 nil
//		predecessor := current.Left
//		for predecessor.Right != current && predecessor.Right != nil {
//			predecessor = predecessor.Right
//		}
//		//d predecessor = findPredecessuor(current
//
//		// 在这里构建了一条指向中序遍历后一个节点的回边
//		if predecessor.Right == nil {
//			predecessor.Right = current
//
//			// 并进行了分支转向
//			current = current.Left
//		} else {
//			predecessor.Right = nil
//			res = append(res, current.Val)
//			current = current.Right
//		}
//	}
//	}
//
//return res
//
//}

func TestInorderTravel(t *testing.T){
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:3}
	r1 := inorderTraversal(root)

	r2 := inorderTraversalWithStack(root)
	r4 := MorriInorderTravel(root)

	r3 := preorderTraversal(root)

	t2 := &TreeNode{Val:1}
	t2.Right = &TreeNode{Val:2}
	r5 := inorderTraversalWithStack(t2)
	t.Log(r1, r2, r3, r4, r5)
}

// 前序遍历的框架思维//
// 根左右
// 如果小于零就不要了
