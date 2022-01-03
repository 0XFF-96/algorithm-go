package main


// * Definition for a Node.
type Node struct {
	     Val int
	     Next *Node
	     Random *Node
}
// 难点：复制链表的同时，维护好链表 random 指针自身的映射关系
// 最容易理解的解题方法是：A->B->C->D。 A->A1->B-B1-C-C1-D-D1

// 算法如下
// 1、for - 在同一个链表内，复制一个新的节点
// 2、for - 复制 Random 指针的相关值
// 3、for - 重新组织相关的函数


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */


// 这里的答案有个致命的错误
// 没有理解清楚题目的意思
// 不能修改原链表的指针
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// copy 链表
	cur := head
	for cur != nil {
		copyNode := &Node{Val:cur.Val}
		next := cur.Next
		cur.Next = copyNode
		copyNode.Next = next

		cur = copyNode.Next
	}

	// copy 链表
	cur = head
	for cur != nil && cur.Next != nil {
		random := cur.Random
		if random == nil {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	dummy := &Node{}
	d := dummy
	for cur != nil && cur.Next != nil {
		d.Next = cur.Next
		d = d.Next
		cur = cur.Next.Next
	}
	return dummy.Next

	// 这是什么错误？
	// Next pointer of node with label 7 from the original list was modified.
}

// 难点：复制链表的同时，维护好链表 random 指针自身的映射关系
// 最容易理解的解题方法是：A->B->C->D。 A->A1->B-B1-C-C1-D-D1
// 算法如下
// 1、for - 在同一个链表内，复制一个新的节点
// 2、for - 复制 Random 指针的相关值
// 3、for - 重新组织相关的函数