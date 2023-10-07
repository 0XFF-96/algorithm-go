package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 非常简洁的代码:
// https://leetcode.com/problems/add-two-numbers/discuss/997/c%2B%2B-Sharing-my-11-line-c%2B%2B-solution-can-someone-make-it-even-more-concise
// Time complexity : O(\max(m, n))O(max(m,n)). Assume that mm and nn represents the length of l1l1 and l2l2 respectively, the algorithm above iterates at most \max(m, n)max(m,n) times.
// 相关代码: https://leetcode.com/problems/add-two-numbers/discuss/1010/Is-this-Algorithm-optimal-or-what
// Space complexity : O(\max(m, n))O(max(m,n)). The length of the new list is at most \max(m,n) + 1max(m,n)+1.
	// 值得一看
// https://leetcode.com/problems/add-two-numbers/solution/
func addTwoNumbers(a *ListNode, b *ListNode) *ListNode {
    tail := &ListNode{0, nil}
    head := tail 
    var num int 
    for a != nil || b != nil {
        num /= 10 
        if a != nil {
            num += a.Val 
            a = a.Next
        }
        if b != nil {
            num += b.Val 
            b = b.Next
        }
        head.Next = &ListNode{Val:num%10}
        head = head.Next 
    }
    
    // 进位
    if num/10 == 1{
        head.Next=&ListNode{Val:1}
    }
    
    return tail.Next
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    for prev,sum:=head,0; l1 != nil || l2 != nil || sum>0; sum/=10 {
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        prev.Next = &ListNode{
            Val : sum%10,
        }
        prev = prev.Next
    }
    return head.Next
}