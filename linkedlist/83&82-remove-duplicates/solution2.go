


// 最终正确的答案
// 从这里，也可以推到出 v1 版的解题方法
// 只要从这个模板里面改一些地方，就可以得到 v1 版的答案
func deleteDuplicatesV2(head *ListNode) *ListNode {
    dummy := &ListNode{}
    d := dummy  
    
    for head != nil  {
        // [1, 1, 1, 3, 3]
		// 应该改为不等与 pre value 才对？
		// 改1
        if isDup(head) {
            for isDup(head) {
                head = head.Next
            }
        } else {
            d.Next = head
            d = d.Next
        }
        // 👇的代码
        head = head.Next
	}
	
	// 如果没有这行代码会怎么样？
	// [1, 2, 3, 4, 5, 5, 5, 5]
    d.Next = nil 
    return dummy.Next
}

func isDup(head *ListNode) bool {
    return head.Next != nil && head.Val == head.Next.Val
}

