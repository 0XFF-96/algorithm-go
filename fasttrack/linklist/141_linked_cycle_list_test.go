package linklist

import (
	"testing"
	"time"
)

// 这种基础题目，一定要一题多解法
// 会成为解决某些难题的模板型题目。
//

// Method: 1、 Hash Table 2、slow、fast pointer
//
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// 看看别人的条件语句是怎么写的：https://leetcode.com/problems/linked-list-cycle/solution/
// 值得学，不同的 idea
// 最最最细节的东西。
//

func TestLog(t *testing.T){
	var (
		in uint64
		build string
	)

	t.Logf("%d %s", in, build)


	now := time.Now()
	tt := now.AddDate(0, 5, 0)
	t.Log(tt.String())
}