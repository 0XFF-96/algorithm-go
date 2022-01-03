package main 

func reorderList(head *ListNode)  {
    if head == nil {
        return 
    }
    stack := []*ListNode{}
    queue := []*ListNode{}
    
    cur := head 
    for cur != nil {
        stack = append(stack, cur)
        queue = append(queue, cur)
        cur = cur.Next
    }
    
    for len(stack) > 0 && len(queue) > 0 {
        sTop := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        qTop := queue[0]
        queue = queue[1:]
        
        // 如果有重复元素怎么办？
        if sTop == qTop || sTop == qTop.Next  {
            sTop.Next = nil 
            break
        }
        sTop.Next = qTop.Next
        qTop.Next = sTop
    }
    return
}

// 感觉利用递归来做，效果会好很多
// 如果是数组的话，会怎么样呢？
// 怎么构建【后向】指针的问题。 
// 
// 首先，想一下。
// 能不能先套路一下？指针的移动
// 
// 最直观的方式，解法1 
// 1、两个 for 循环
// 2、第二个 for 循环，用户寻找倒数第二个节点。
//
// 解法2
// 利用 queue 和 stack 
// 空间换时间的做法。
// 