

// https://leetcode.com/problems/swap-nodes-in-pairs/discuss/149435/golang-iterative-solution-0ms
// 
// 
// public class Solution {
//     public ListNode swapPairs(ListNode head) {
//         if ((head == null)||(head.next == null))
//             return head;
//         ListNode n = head.next;
//         head.next = swapPairs(head.next.next);
//         n.next = head;
//         return n;
//     }
// }

// 有两种做法
// 1、递归
// 2、

// __topic__: daycam-metric-prod and attr.uri = "notification-guide-clicked" | SELECT   timestamp - (timestamp + 8 * 3600) % 86400 AS time, count( distinct user_id ) AS c, '每日点击人数' as act where "attr.device" = 'oppo' and "attr.os_version" >= 29  group by time
// s


func swapPairs(head *ListNode) *ListNode {
    // 为什么是 || 号
    if head == nil || head.Next == nil {
        return head
    }
    
    rec := head.Next 
    
    // 为什么填 .Next.Next
    head.Next = swapPairs(head.Next.Next)
    rec.Next = head
    return rec
}