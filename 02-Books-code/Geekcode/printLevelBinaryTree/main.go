package printLevelBinaryTree

import (
	"testing"
)

// 重建二叉树
// https://www.nowcoder.com/practice/8a19cbe657394eeaac2f6ea9b0f6fcf6?tpId=13&tqId=11157&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking
// 输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字
// 需要递归遍历，找到节点
// 前：左根右， 中: 根左右
type TreeNode struct{
	val int
	right *TreeNode
	left *TreeNode
}


// 从上打印二叉树
// 栈操作
// dfs 算法遍历
func TestPrintLevelBinaryTree(t *testing.T){
	// l := list.New() 尝试一下是否能够
	//
	tree := &TreeNode{val:1}
	tree.left = &TreeNode{val:2}
	tree.right = &TreeNode{val:3}
	tree.right.left = &TreeNode{val:4}

	res := printBinaryTreeFromTop(tree)
	t.Logf("%v", res)
}

func printBinaryTreeFromTop(tree *TreeNode) []int {
	if tree == nil {
		return nil
	}
	var res []int
	var queue []*TreeNode
	queue = append(queue, tree)
	for;len(queue) > 0 ;{
		node := queue[0]
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		queue = queue[1:]
		res = append(res,node.val)
	}
	return res
}
