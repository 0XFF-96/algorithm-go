package main

import (
	"strconv"
)

type Stack []int
func (s *Stack) Empty() bool {return len(*s) == 0}
func (s *Stack) Push(x int) {*s = append(*s, x)}
func (s *Stack) Peek() int {return (*s)[len(*s)-1]}
func (s *Stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

type MyQueue struct {
	s1, s2 Stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int)  {
	q.s1.Push(x)
}


// 如何没有任何元素，然后执行了 pop 会怎么办？
//
/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if q.s2.Empty() {
		for !q.s1.Empty() {
			x := q.s1.Pop()
			q.s2.Push(x)
		}
	}
	return q.s2.Pop()
}


/** Get the front element. */
func (q *MyQueue) Peek() int {
	if q.s2.Empty() {
		for !q.s1.Empty() {
			x := q.s1.Pop()
			q.s2.Push(x)
		}
	}
	return q.s2.Peek()
}


/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.s1.Empty() && q.s2.Empty()
	strconv.Atoi()
}

