package reverse_link

import (
	"fmt"
	"testing"
)

// 从尾部打印链表
type LinkNode struct {
	val int
	next *LinkNode
}

func makeLinkFromSlice(s []int) *LinkNode{
	if len(s) == 0 {
		return nil
	}
	head := &LinkNode{val:s[0],next:nil}
	dummy := head
	for _, val := range s[1:]{
		node := &LinkNode{val:val,next:nil}
		dummy.next = node
		dummy = node
	}
	return head
}

func TestReverseLink(t *testing.T){
	link := makeLinkFromSlice([]int{1, 2, 4, 5, 6})

	// 如何处理边界条件？
	// reverseLink 的终止条件是什么？
	// 能不能写一个 递归的 reverseLink 的写法？
	reLink := reverseLink(link)
	// t.Log(reLink)
	printLink(reLink)
}


func printLink(link *LinkNode){
	for ; link != nil; {
		fmt.Println(link.val)
		link = link.next
	}
}

func reverseLink(link *LinkNode) *LinkNode {
	if link == nil {
		return nil
	}
	var pre *LinkNode
	head := link
	next := link
	for ; head != nil;  {
		next = next.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}
