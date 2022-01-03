package main

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */


type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	head   *Node
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/// ** Get the value of the index-th node in the linked list.
// If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {

	if index < 0 || index > this.length-1 {
		return -1
	}

	cur := this.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}

	if cur == nil {
		return -1
	}

	return cur.Val

}

// ** Add a node of value val before the first element of the linked list.
// After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {

	var newNode Node
	newNode.Val = val
	newNode.Next = this.head
	this.head = &newNode

	this.length++

}

//  Append a node of value val to
//  the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
		return
	}

	cur := this.head
	for i := 0; cur.Next != nil; i++ {
		cur = cur.Next
	}

	var newNode Node
	newNode.Val = val
	cur.Next = &newNode

	this.length++
}

//  Add a node of value val before the index-th node in the linked list.
//  If index equals to the length of linked list,
//  the node will be appended to the end of linked list.
//  If index is greater than the length, the node will not be inserted.
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	} else if index == this.length {
		this.AddAtTail(val)
		return
	} else if index < 0 || index > this.length {
		return
	}

	cur := this.head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}

	var newNode Node
	newNode.Val = val
	newNode.Next = cur.Next
	cur.Next = &newNode

	this.length++

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	} else if index == 0 {
		this.head = this.head.Next
		this.length--
		return
	}

	cur := this.head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next
	this.length--

}

