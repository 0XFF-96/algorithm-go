package main

import (
	"fmt"
)

type ListNode struct {
	     Val int
	     Next *ListNode
}


// 有❌，不 AC
// 但是，思路是类似的，差不多都能够解决问题。
//

// edge: 判断条件写错❌
// [0]
//
func numComponents(head *ListNode, G []int) int {
	if len(G) == 0 {
		return 0
	}

	// unique integer values，提示很明显需要用 set 来做
	set := map[int]struct{}{}

	// 这里...
	// 应该是把 G 的元素放进去
	// 而不是把 listNode 的元素放进去
	// 错误1
	cur := head
	for i:=0; i < len(G); i++ {
		set[G[i]] = struct{}{}
	}

	cur = head
	count := 0
	for cur != nil && cur.Next != nil {
		_, ok1 := set[cur.Val]
		_, ok2 := set[cur.Next.Val]

		fmt.Println(cur.Val)
		// 第一个节点在里面
		// 第二个节点不在里面
		// 重复加了？
		if (ok1 && !ok2) || (!ok1 && ok2 ) {
			count++
		}
		cur = cur.Next
	}

	return count
	// 连通分量
	// connected components
	// split it
	//
	// 用一个 set 记录 G []int
	// 然后，遍历的时候，看看这个节点有没有在 set 里面？
	// 怎么区分，本来就不在 set 里面

	// 怎么表达 connect 关系？
	// dfs
}


// 相关情况
// Plus One Linked List
// Flatten a Multilevel Doubly Linked List
// Insert into a Sorted Circular Linked List
func numComponentsV1(head *ListNode, G []int) int {
	var ans int

	var gMap = make(map[int]bool, len(G))
	for i := range G {
		gMap[G[i]] = true
	}

	for curr, i := head, 0; curr != nil; curr, i = curr.Next, i+1 {
		if _, ok := gMap[curr.Val]; !ok {
			continue
		}

		if curr.Next == nil {
			ans++
			continue
		}

		if _, ok := gMap[curr.Next.Val]; !ok {
			ans++
		}
	}

	return ans
}