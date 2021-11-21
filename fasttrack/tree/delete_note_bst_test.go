package tree

import (
	"testing"
)
// BST 的定义
// less than or equal to the node's key
// The right subtree of a node contains only nodes with keys greater than or equal to the node's key.

func deleteNodeWrong(root *TreeNode, key int) *TreeNode {

	deletedNode,_ := searchBSTNode(root, key)
	// need to know its parent.
	// parent :=
	if deletedNode.Left != nil && deletedNode.Right != nil {

		return nil
	} else {


		return nil
	}
}

func TestSearchNode(t *testing.T){
	tree1 := &TreeNode{Val:1}

	tree2 := &TreeNode{Val:2}

	//Tree3 := &TreeNode{Val:3}

	bst := &TreeNode{Val:0}
	bst.Left = tree1
	bst.Right = tree2
	// bst.Right.Right

	// 这个函数还不够健壮
	p,c := searchBSTNode(tree1, 2)
	t.Log(p, c)
}


// 找到要删除的节点和它的父亲节点
//
func searchBSTNode(root *TreeNode, key int) (parent*TreeNode,child*TreeNode) {
	if root == nil || root.Left == nil && root.Right == nil {
		return nil,nil
	}
	if root.Left != nil && root.Left.Val == key  {
		return root, root.Left
	} else if root.Right != nil && root.Right.Val == key {
		return root, root.Right
	}

	if root.Left != nil && root.Left.Val > key {
		return searchBSTNode(root.Left, key)
	} else if root.Right != nil {
		return searchBSTNode(root.Right, key)
	}
	// 左右孩子都为空的情况
	return nil, nil
}


// 来自于这篇文章：https://leetcode.com/problems/delete-node-in-a-bst/discuss/93296/Recursive-Easy-to-Understand-Java-Solution
// Recursively find the node that has the same value as the key, while setting the left/right nodes equal to the returned subtree
// Once the node is found, have to handle the below 4 cases
// node doesn't have left or right - return null
// node only has left subtree- return the left subtree
// node only has right subtree- return the right subtree
// node has both left and right - find the minimum value in the right subtree, set that value to the currently found node, then recursively delete the minimum value in the right subtree
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Right.Val > key {
		root.Right = deleteNode(root.Right, key)
	} else {
		// 找到需要 delete 的节点了
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// 左右孩子都存在都情况
		min := findMin(root.Right)
		root.Val = min.Val
		root.Right = deleteNode(root.Right, min.Val)
	}
	return root
}


func findMin(root *TreeNode)*TreeNode{
	for root.Left != nil {
		root = root.Left
	}
	return root
}



// 尝试一下非递归解法
// https://leetcode.com/problems/delete-node-in-a-bst/discuss/93298/Iterative-solution-in-Java-O(h)-time-and-O(1)-space
//



// 这个版本的 Go Solution 太繁琐了
// https://leetcode.com/problems/delete-node-in-a-bst/discuss/93381/A-golang-solution


// 这版本，将 adjust tree 的部分分离开来了感觉不错
// https://leetcode.com/problems/delete-node-in-a-bst/discuss/147665/Golang-solution

