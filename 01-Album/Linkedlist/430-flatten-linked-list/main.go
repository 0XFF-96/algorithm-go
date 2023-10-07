package main

// 这个数据结构
// 适用于什么场景？
// 能用来干什么呢？
type Node struct {
	     Val int
	     Prev *Node
	     Next *Node
	     Child *Node
}

func flatten(root *Node) *Node {
	// To serialize all levels together we will add nulls in each level to signify no node
	// connects to the upper node of the previous level. The serialization becomes
	//
	// 1. Merging the serialization of each level
	// 2. removing trailing nulls
	if root == nil {return nil}
	prev := &Node{Val:-1}
	stack := []*Node{root}
	for len(stack) > 0 {
		curr := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		curr.Prev = prev
		prev.Next = curr
		prev = curr
		if curr.Next != nil {
			stack = append(stack, curr.Next)
		}
		if curr.Child != nil {
			stack = append(stack, curr.Child)
			curr.Child = nil
		}
	}
	root.Prev = nil
	return root
}


