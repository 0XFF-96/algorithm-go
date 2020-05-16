package main

import (
	"testing"
)

func TestT(t *testing.T){
	node := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:2,
			Next: &ListNode{
				Val: 3,
		 		Next: &ListNode{
					Val:1,
				},
			},
		},
	}
	insertionSortListV2(node)
}
