package tree

import (
	"testing"
)

// 脑海中要想像有一个图形
// 难点在构建，指向下一个节点的回边
//
// https://www.youtube.com/watch?v=wGXB9OWhPTg
func MorriInorderTravel(root *TreeNode)[]int{
	var res []int
	current := root
	for current != nil {
		if current.Left == nil {
			res = append(res, current.Val)
			current = current.Right
		} else {
			// find the predecessor.
			// To Find the predecess keep going right
			// till right node is not null or right node
			// is not current.

			// 这里找 predecessor 的前提是，左分支不为 nil
			predecessor := current.Left
			for predecessor.Right != current && predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			//d predecessor = findPredecessuor(current

			// 在这里构建了一条指向中序遍历后一个节点的回边
			if predecessor.Right == nil {
				predecessor.Right = current

				// 并进行了分支转向
				current = current.Left
			} else {
				predecessor.Right = nil
				res = append(res, current.Val)
				current = current.Right
			}
		}
	}

	return res

}


// 难点
// 1、构建和删除临时回边
// 2、在有假设条件下，找到前继节点


// 2.10 尝试独立写出来，还是失败了
func MorriInorderTravel2(root *TreeNode)[]int {
	//if root ==  nil {
	//	return nil
	//}
	//
	//var ret []int
	//cur := root
	//var predecessor *TreeNode
	//
	//for cur != nil {
	//	if cur.Left == nil {
	//
	//
	//
	//	} else {
	//		// 构建回边
	//		right := cur.Left
	//		for right.Right != nil {
	//			right = right.Right
	//		}
	//		predecessor = right
	//
	//		cur = cur.Left
	//	}
	//
	//
	//}
	return nil
}

// 用 Morris 的方法作为基础，构建后向边
// 引用
// https://leetcode.com/problems/binary-tree-preorder-traversal/discuss/45281/Golang-Morris-Traversal-(using-constant-space)-and-iterative-solution-(using-stack)
// https://leetcode.com/problems/binary-tree-preorder-traversal/discuss/45281/Golang-Morris-Traversal-(using-constant-space)-and-iterative-solution-(using-stack)
// 这个 inorder 更强
// https://leetcode.com/problems/binary-tree-inorder-traversal/discuss/31282/golang-2-solutions-morris-traversal-iterative-using-a-queuehttps://leetcode.com/problems/binary-tree-inorder-traversal/discuss/31282/golang-2-solutions-morris-traversal-iterative-using-a-queue
// https://www.youtube.com/watch?v=wGXB9OWhPTg
// https://leetcode.com/problems/binary-tree-inorder-traversal/discuss/31282/golang-2-solutions-morris-traversal-iterative-using-a-queue


func TestMorii(t *testing.T){
	tree1 := initTree()

	ret1 := MorriInorderTravel(tree1)
	ret := MorriInorderTravel3(tree1)
	t.Log(ret)
	t.Log(ret1)
}


func MorriInorderTravel3(root *TreeNode)[]int {
	if root == nil {
		return nil
	}

	cur := root
	var predecessor *TreeNode
	var ret []int

	for cur != nil {
		if cur.Left == nil {
			ret = append(ret, cur.Val)
			cur = cur.Right
		} else {
			predecessor = cur.Left
			// 这里面的条件写错了

			// 一定是 prodecessor.Right 作为条件判断.
			for predecessor.Right != cur && predecessor.Right != nil {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				predecessor.Right = cur

				// 前序遍历
				cur = cur.Left
			} else {
				predecessor.Right = nil
				// 中序遍历
				ret = append(ret, cur.Val)
				cur = cur.Right
			}
		}
	}
	return ret
}
