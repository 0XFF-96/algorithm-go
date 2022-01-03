package main

func reorderListV2(head *ListNode)  {
    cnt := 0
    cur := head
    
    for cur != nil {
        cnt++
        cur = cur.Next
    }
    
    // revert list from (n+1)/2 +1
    
    cur = head
    var pre *ListNode
    for i := 0; i < (cnt+1)/2; i++ {
        pre = cur
        cur = cur.Next
    }
    
    if pre != nil {
        pre.Next = nil
    }
    
    head2 := revert(nil, cur)
    nilHead := &ListNode{}
    cur = nilHead
    
    for head != nil {
        cur.Next = head
        head = head.Next
        cur = cur.Next
        
        if head2 != nil {
            cur.Next = head2
            head2 = head2.Next
            cur = cur.Next
        }
    }
}

func revert(pre, cur *ListNode) *ListNode {
    if cur == nil {
        return pre
    }
    next := cur.Next
    cur.Next = pre
    return revert(cur, next)
}