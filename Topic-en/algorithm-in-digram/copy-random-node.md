
- 复杂链表的复制

```

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil 
    }

    m := map[*Node]*Node{} // 建立 old and new 的 mapping 关系
    cur := head 
    for cur != nil {
        m[cur] = &Node{Val: cur.Val}
        cur = cur.Next 
    }

    // 复制 random 链表 
    cur = head 
    for cur != nil {
        newNode := m[cur]
        newNode.Next = m[cur.Next]
        newNode.Random = m[cur.Random]
        cur = cur.Next
    }
    
    // return head of new linklist node 
    return m[head]
}

```
