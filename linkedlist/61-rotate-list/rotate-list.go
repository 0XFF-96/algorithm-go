

// fast 
// slow 
// 把链表连起来，然后 break 操作。
// 
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil 
        
    }
    
    if head.Next == nil {
        return head
    }
    
    fast := head 
    slow := head
    
    count := 0 
    dummy := head 
    
    for dummy != nil {
        dummy = dummy.Next 
        count++
    }
    
    k = k % count
    
    if k == 0 {
        return head 
    }
    // k is non-negative 
    // so k may be overflow ?
    // how to deal with it ?
    i := 0
    for i < k && fast != nil {
        fast = fast.Next
        i++
    }
    
    if fast == nil {
        return head
    }
    
    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }
    
    // fmt.Println(fast.Val)
    // fmt.Println(slow.Val)
    // 
    curHead := slow.Next
    slow.Next = nil 
    fast.Next = head 
    return curHead
}

// 1、
// [1,2,3,4,5]
// 2
// [4,5,1,2,3]
// 这几种情况怎么办？
// [1] 0 、 [1, 2] 2