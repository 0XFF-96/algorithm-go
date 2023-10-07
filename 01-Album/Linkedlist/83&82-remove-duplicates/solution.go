
// 错误的做法。
// 不合理的

// 正确代码。
func deleteDuplicates(head *ListNode) *ListNode {
    cur := head 
    for cur != nil && cur.Next != nil {
        if cur.Next.Val == cur.Val {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    return head
}


func deleteDuplicates(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head 
    dummy := pre
    for cur != nil {
        if pre == nil  {
            pre = cur
            cur = cur.Next 
        }
        if pre.Val = cur.Val {
            cur.Next = cur.Next.Next 
        }

    }
    return 
}
