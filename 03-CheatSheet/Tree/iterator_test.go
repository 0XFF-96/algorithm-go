package tree

import (
	"testing"
)

// 接口与实现编程题目
// 优雅的接口封装
// 中序遍历的下一个节点
// 如何记录中序遍历的值？
//

func TestBSTIterator(t *testing.T){
	tree := initTree()


	iterator := Constructor(tree)
	t.Log(iterator.next.cur.Val)
	t.Log(iterator.next.next.cur.Val)
}


type BSTIterator struct {
	cur *TreeNode
	next *BSTIterator
}

// 中序遍历的下一个节点
func Constructor(root *TreeNode) BSTIterator{
	if root == nil {
		return BSTIterator{nil, nil }
	}
	stack := []*TreeNode{}
	cur := root

	ret := &BSTIterator{nil, nil}
	pre := ret
	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[0:len(stack)-1]

			// 连接下一个节点
			pre.next = &BSTIterator{cur:cur}
			pre = pre.next
			cur = cur.Right
			continue
		}

		stack = append(stack, cur)
		cur = cur.Left
	}

	return *ret
}

func (this *BSTIterator) Next() int {
	if this.next == nil {
		return 0
	}

	val := this.next.cur.Val
	// nextNode := this.next.cur
	this.cur = this.next.cur
	this.next = this.next.next

	// 错误点在这个地方
	// 反思一下，为什么会错在这里？
	// 	this = this.next
	// 为什么这句赋值语句是错误的？
	// 感觉是 this 是指这个对象的地址，是在运行时是不能修改的...
	return val
}

func (this *BSTIterator) HasNext() bool {
	if this.next != nil {
		return true
	}

	return false
}

// https://www.youtube.com/watch?v=nzmtCFNae9k

// 高级点的方法
// https://leetcode.com/problems/binary-search-tree-iterator/discuss/352302/golang-morris-traversal
// let it go
//


// 视频里面的另一种做法，是使用栈。
//