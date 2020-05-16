

# 这是一个 follow-up 问题
1-https://leetcode.com/problems/reverse-linked-list-ii/solution/ 【非常重要的题目】

2-这个 youtu video 很好地解析了，recursive list 的相关作用。https://www.youtube.com/watch?v=O0By4Zq0OFc

3- https://leetcode.com/problems/reverse-linked-list-ii/solution/ 
4- 




### 参考这个人的方法做出来的
- 他的整个思考过程非常清晰
- 思路也一直很好，步步递进。 
- 从分析【数组】的反向遍历开始。 
- 如果数据结构是【数组】而不是【链表】？该怎么办？
```
To solve the problem recursively, step by step.
First, I know the classic recursive way to reverse a linked list:

    ListNode reverse(ListNode head) {
        if (head.next == null) return head;
        ListNode last = reverse(head.next);
        head.next.next = head;
        head.next = null;
        return last;
    }
Then I come up this way to reverse the first N elements:

    ListNode successor = null;
    ListNode reverseN(ListNode head, int n) {
        if (n == 1) {
            successor = head.next;
            return head;
        }
        ListNode last = reverseN(head.next, n - 1);
        head.next.next = head;
        head.next = successor;
        return last;
    }    
Finally I solve this problem:

    public ListNode reverseBetween(ListNode head, int m, int n) {
        if (m == 1) {
            // You can also expand the code here to get rid of the helper function 'reverseN'
            return reverseN(head, n);
        }
        head.next = reverseBetween(head.next, m - 1, n - 1);
        return head;
    }
```    