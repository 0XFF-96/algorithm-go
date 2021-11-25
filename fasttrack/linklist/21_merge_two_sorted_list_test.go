package linklist

import (
	"testing"
)

// 如果直接用 goland 带的 run, 会出现报错
// # command-line-arguments [command-line-arguments.test]

// 因为 setup call 有问题？
// GOROOT=/Users/lijinrui/go/go1.13.5 #gosetup
// GOPATH=/Users/lijinrui/go #gosetup
// /Users/lijinrui/go/go1.13.5/bin/go test -c -o /private/var/folders/09/r6f6jypj0f9bj7rz2qpv6jyr0000gn/T/___TestMerge_in_21_merge_two_sorted_list_test_go /Users/lijinrui/algorithm-go/fasttrack/linklist/21_merge_two_sorted_list_test.go #gosetup

func generateLink(values []int) *ListNode {
	var head *ListNode
	cur := head
	for _, v := range values {
		if cur == nil {
			cur = &ListNode{}
		}
		cur.Val = v
		cur = cur.Next
	}
	return head
}

func TestMerge(t *testing.T) {
	l1 := generateLink([]int{1, 23, 3})
	l2 := generateLink([]int{4, 5, 6})
	ret := mergeTwoLists(l1, l2)
	t.Log(ret)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{Val:0}
	cur := dummyNode
	for l1 != nil && l2 != nil {
		head1 := l1.Val
		head2 := l2.Val
		if head1 <= head2 {
			cur.Next = l1
			l1 = l1.Next  // 这里会不会影响上层的 l1 的相关指针？
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}
	return dummyNode.Next
}
