package main

type ListNode struct {
	Next *ListNode
}
// 根据别人思路写出来的代码
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    // reverse 一个 array 是怎么样的？
    //  o(n) o(n) 复杂度算法
    
    // O(n) O(1) 复杂度？
    
    // So, we rely on recursion to simulate the backward pointer
    // 利用递归，可以构造单链表的 backward 指针。
    var reverseN func(head *ListNode, n int) *ListNode 
    
    var successor *ListNode
    reverseN = func(head *ListNode, n int) *ListNode {
        if n == 1 {
            successor = head.Next 
            return head 
        }
        
        last := reverseN(head.Next, n-1)
        head.Next.Next = head 
        head.Next = successor 
        return last
    }
    
    if m == 1 {
        return reverseN(head, n)
    }
    
    head.Next = reverseBetween(head.Next, m-1, n-1)
    return head 
}