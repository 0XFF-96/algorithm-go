

- 反转链表

```

func reverseList(head *ListNode) *ListNode {
    cur := head
    var pre *ListNode
    for cur != nil {
        tmp := cur.Next  // 暂存后继节点
        cur.Next = pre   // 修改 next 引用指向 
        pre = cur        // pre 暂存 cur 
        cur = tmp        // cur 访问下一节点
    }
    return pre 
}

```

- （重点）

```

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head 
    }

    tail := reverseList(head.Next) // 1. head.Next 继续
    head.Next.Next = head // 2. 反转
    head.Next = nil // 3. 
    // 1 -> 2 -> 3 
    // 1
    // 2. *    3 -> 2 
    // 3
    return tail
}

```


- 递归写法2 

```

- 方法 3 
```

func reverseList(head *ListNode) *ListNode {
    return rev(head, nil)
}


func rev(cur, pre *ListNode) *ListNode {
    // 1. 终止递归
    if cur == nil {
        return pre 
    }

    res := rev(cur.Next, cur)  // 递归修改后继节点
    cur.Next = pre            //  修改节点引用指向 
    return res                //  返回反转链表头节点
}

```

- 错误写法
- 为什么， 错在哪里？
```

func reverseList(head *ListNode) *ListNode {
    if head == nil  {
        return head 
    }

    tail := reverseList(head)

    head.Next = head // 这一步死循环了
    // tail := reverseList(head.Next) // 1. head.Next 继续
    // head.Next.Next = head // 反转
    // head.Next = nil // 差了一步
    // 1 -> 2 -> 3 
    // 1
    // 2. *    3 -> 2 
    // 3
    return tail
}

```
