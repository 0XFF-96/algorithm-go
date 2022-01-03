package main
// 不 ac 的做法.
// 
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // could you do it in one-pass ?
    if head == nil {
        return nil 
    }
    
    fast := head 
    slow := head 
    
    // n-th
    
    i := 0 
    for fast != nil {
        fast = fast.Next
        if i == n {
            break 
        }
        i++
    }
    if fast == nil {
        return nil 
    }
    
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    
    slow.Next = slow.Next.Next
    return head
}


// 正确的 vc

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // could you do it in one-pass ?
    if head == nil {
        return nil 
    }
    
    dummy := &ListNode{}
    dummy.Next = head
    
    fast := dummy
    slow := dummy
    
    // n-th
    
    i := 0
    for fast != nil && i != n{
        fast = fast.Next
        i++
    }
    
    if fast == nil {
        return nil
    }
    
    // fast.Next 不等于 0 才行
    // two point 的相关概念
    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }
    
    slow.Next = slow.Next.Next
    return dummy.Next
}
