

func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil 
    }
    // 是否循环
    // 循环的节点在哪里？
    
    // 当两个节点相遇当时候，
    // 另外一个节点，应该从哪里重启
    // 计算出数学公式
    
    slow := head 
    fast := head 
    
    
    // 这里填错了
    // slow != nil && fast.Next != nil 
    // slow != nil || fast.Next != nil 
    // 这几个讯号之间有什么差别？
    // 
    for fast != nil && fast.Next != nil {
        // slow = slow.Next 
        // fast = fast.Next.Next 
        // if slow == fast {
           // break 
        // }
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }
    
    if fast == nil || fast.Next == nil {
        return nil
    }
    
    // 重启
    fast = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    
    return slow
}