
### delete link note 


````

func deleteNode(head *ListNode, val int) *ListNode {
    if head == nil{
        return nil
    }
    dummy := &ListNode{Next:head}
    fast,slow := head, dummy
    for fast != nil{
        if fast.Val == val{
            slow.Next = fast.Next
            break
        }
        fast,slow = fast.Next, slow.Next
    }
    return dummy.Next
}

```
