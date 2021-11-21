package tree

import (
	"fmt"
	"testing"
)

// Populating Next Right Pointers in Each Node
// 用层次遍历的框架可以解决这个问题

type Node struct {
	val int
	left *Node
	right *Node
	next *Node
}
func connect(root *Node){
	if root == nil {
		return
	}

	// curRoot := root
	queue := []*Node{root}
	for len(queue) != 0 {
		size := len(queue)
		for i:=0; i < size; i++{
			// dequeue
			curNode := queue[0]
			fmt.Println(curNode.val)
			queue = queue[1:]

			// 加入噶时机唔对...
			// 要加一个 pre 变量
			if len(queue) > 1 {
				fmt.Println("-----")
				curNode.next = queue[1]
				fmt.Println(queue[1].val)
				fmt.Println("------")
			}

			if curNode.left != nil {
				queue = append(queue, curNode.left)
			}

			if curNode.right != nil {
				queue = append(queue, curNode.right)
			}
		}

	}
}

func TestConnect(t *testing.T){
	root := &Node{val:1}
	root.right = &Node{val:2}
	root.left = &Node{val:3}
	root.left.left = &Node{val:4}
	root.left.right = &Node{val:5}

	connect(root)
	t.Log(root.next)
	//t.Log(root.left.next.val)

	// errors.Is()
}


// python 递归解法
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/discuss/37715/Python-solutions-(Recursively-BFS%2Bqueue-DFS%2Bstack)
//


// 用完全二叉树，这种做是.. 减轻工作量的一个基本方式
//