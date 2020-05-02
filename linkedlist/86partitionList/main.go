func partition(head *ListNode, x int) *ListNode {
    less := &ListNode{}
    greater := &ListNode{}
    cur := head 
    lessD := less
    greaterD := greater
    
    for cur != nil {
        if cur.Val >= x {
            greaterD.Next = cur
            greaterD = greaterD.Next
        } else {
            lessD.Next = cur 
            lessD = lessD.Next
        }
        cur = cur.Next
	}
	// 这里少了一行代码，
	// 有个坑
    lessD.Next = greater.Next
    
    return less.Next
}
