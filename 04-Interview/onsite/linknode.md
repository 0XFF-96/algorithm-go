

```

package main

import (
    "fmt"
)
func main() {
    //a := 0
    //fmt.Scan(&a)
    //fmt.Printf("%d\n", a)
    fmt.Printf("Hello World!\n");
}

type LinkNode struct {
    Val int 
    Next *LinkNode
}
// 输入: k = 3, 1 -> 2 -> 3 -> 4 -> 5 
// 3 - > 4 -> 5 - > 1 -> 2

// 思路
// k = 3 
// 1. 1 -> 2 -> 3 -> 4 -> 5 , 找到 3        
// 2. 断开，tmp1-> 1 -> 2   tmp -> 3 -> 4 -> 5(rear)-> nil
// 3. 合并, rear -> tmp1,  
//   dummyNode -> 3 -> 4 -> 5-> 1 -> 2

// k = 7, 1 -> 2 -> 3 -> 4 -> 5 
// 输出， 

func reverseKLinkNode(node *LinkNode, k int) *LinkNode {
    if k == 0 {
        return node
    }

    dummyNode := &LinkNode{}
    dummyNode.Next = node 

    // 1.
    // k := k % len(node)
    // k < 链表的长度 ， 拿到链表长度，取模。
    rear := dummyNode.Next 
    count := 1
    for rear != nil && rear.Next != nil {
        rear = rear.Next 
        count += 1 
    }

    k = k % count // 取模，保证不会溢出
    newHeaderNode := dummyNode.Next
    for i := 0; i <= k; i ++ {
        // 断开链表
        if i == k {
            tmp2 := newHeaderNode.Next
            newHeaderNode.Next = nil 
            newHeaderNode = tmp2
            break 
        }
        newHeaderNode = newHeaderNode.Next // 2
    }

    // 3. 把原来的尾指针，指向原来的头节点
    // 组装两个链表
    rear.Next = dummyNode.Next
    // 返回新的头节点
    return newHeaderNode
}

```
