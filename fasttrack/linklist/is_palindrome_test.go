package linklist

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T){
	l := ListNode{Val:1}
	l.Next = &ListNode{Val:1}
	l.Next.Next = &ListNode{Val:2}
	l.Next.Next.Next = &ListNode{Val:1}

	flag := isPalindrome(&l)
	t.Log(flag)
	// 这种函数是有副作用的。
	// 会修改原来的节点.. 所有
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	// 找出 link list 的中点
	// 存为一个数组
	//
	// could you do it in O(n) time O(1) space?
	// reserve list
	// and compare it with original
	// 2n + O(1)
	// headv2 := head
	rev := reverseListV2(head)
	_ = rev
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
	//for headv2 != nil && rev != nil {
	//	if head.Val != rev.Val {
	//		return false
	//	}
	//	fmt.Println(rev.Val)
	//	headv2 = headv2.Next
	//	rev = rev.Next
	//}

	// for rev != nil {
	//     fmt.Println(rev)
	//     rev = rev.Next
	// }
	return true
}

// 利用栈的特性
// slow, fast 指针
// n/2 空间 O(n) 时间
func isPalindromeV3(head *ListNode) bool {
	fast := head
	slow := head
	stack := []int{}

	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != stack[len(stack)-1] { return false }
		slow = slow.Next
		stack = stack[:len(stack)-1]
	}
	return true
}

func isPalindromeV2(head *ListNode) bool {
	var l []int
	for head != nil {
		l = append(l,head.Val)
		head = head.Next
	}

	if len(l) < 2 {
		return true
	}

	for i := range l {
		if l[i] != l[len(l)-1-i] {
			return false
		}
	}
	return true
}