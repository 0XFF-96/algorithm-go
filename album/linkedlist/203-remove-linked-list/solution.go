package main 


func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{}
    curDummy := dummy
    cur := head
    
    // 应该从 Next 开始判断？
    // 还是从当前节点判断？
    for cur != nil {
        if cur.Val == val {
            cur = cur.Next
            continue 
        }
        curDummy.Next = cur 
        curDummy = curDummy.Next
        cur = cur.Next
    }
    return dummy.Next
}