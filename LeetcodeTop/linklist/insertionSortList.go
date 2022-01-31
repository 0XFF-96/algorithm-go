

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    dummyNode := &ListNode{}
    dummyNode.Next = head 
    lastSortedNode := head 
    curNode := head.Next

    for curNode != nil {
        if lastSortedNode.Val <= curNode.Val {
            lastSortedNode = lastSortedNode.Next
        } else {
            // 从头开始寻找 ，寻找插入位置
            prev := dummyNode
            for prev.Next.Val <= curNode.Val {
                prev = prev.Next
            }
            // 找到了
            // 
            lastSortedNode.Next = curNode.Next
            curNode.Next = prev.Next
            prev.Next = curNode
        }
        // 更新节点
        curNode = lastSortedNode.Next
    }

    return dummyNode.Next
}
