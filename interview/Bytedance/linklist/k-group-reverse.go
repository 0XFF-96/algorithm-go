


```
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序


你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换
```

// 思路

// ## 思路:

// **思路一:**

// 用栈,我们把k个数压入栈中,然后弹出来的顺序就是翻转的!

// 这里要注意几个问题

// 第一,剩下的链表个数够不够k个(因为不够k个不用翻转);

// 第二,已经翻转的部分要与剩下链表连接起来

// **思路二:**

// 尾插法

// 直接举个例子: k = 3


// 递归和迭代两种解法


// 详细步骤

// 链表分区为已翻转部分+待翻转部分+未翻转部分
// 每次翻转前，要确定翻转链表的范围，这个必须通过 k 此循环来确定
// 需记录翻转链表前驱和后继，方便翻转完成后把已翻转部分和未翻转部分连接起来
// 初始需要两个变量 pre 和 end，pre 代表待翻转链表的前驱，end 代表待翻转链表的末尾
// 经过k此循环，end 到达末尾，记录待翻转链表的后继 next = end.next
// 翻转链表，然后将三部分链表连接起来，然后重置 pre 和 end 指针，然后进入下一次循环
// 特殊情况，当翻转部分长度不足 k 时，在定位 end 完成后，end==null，已经到达末尾，说明题目已完成，直接返回即可
// 时间复杂度为 O(n*K)O(n∗K) 最好的情况为 O(n)O(n) 最差的情况未 O(n^2)O(n 
// 2
//  )
// 空间复杂度为 O(1)O(1) 除了几个必须的节点指针外，我们并没有占用其他空间


// 比较清晰的做法，但是变量名字不够清晰。
// public ListNode reverseKGroup(ListNode head, int k) {
//     ListNode dummy = new ListNode(0);
//     dummy.next = head;

//     ListNode pre = dummy;
//     ListNode end = dummy;

//     while (end.next != null) {
//         for (int i = 0; i < k && end != null; i++) end = end.next;
//         if (end == null) break;
//         ListNode start = pre.next;
//         ListNode next = end.next;
//         end.next = null;
//         pre.next = reverse(start);
//         start.next = next;
//         pre = start;

//         end = pre;
//     }
//     return dummy.next;
// }


func reverseKGroup(head *ListNode, k int) *ListNode {
    hair := &ListNode{Next: head}
    pre := hair

    for head != nil {
        tail := pre
        for i := 0; i < k; i++ {
            tail = tail.Next
            if tail == nil {
                return hair.Next
            }
        }
        nex := tail.Next
        head, tail = myReverse(head, tail)
        pre.Next = head
        tail.Next = nex
        pre = tail
        head = tail.Next
    }
    return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
    prev := tail.Next
    p := head
    for prev != tail {
        nex := p.Next
        p.Next = prev
        prev = p
        p = nex
    }
    return tail, head
}