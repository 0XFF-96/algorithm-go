

// 回文链表

// 算法流程， （不考虑链表复原）？


// 思路1：


// 递归+后序遍历 事实上就相当于一个栈，把链表节点值都入栈了，
// 在最开始的时候就用一个变量把头节点存起来，
// 当后序遍历到最后一个节点的时候，就依次比较left节点和right节点即可

// 1 2 3 4 3 2 1
// 1. 这种做法，是每个数字都会被判定两次的。 
// 2. 操作系统的栈 pop ，是比较耗费 cpu 和内存的。 
// 3. 比较清晰，思路可以。 

// 左侧指针
let left;
var isPalindrome = function (head) {
  left = head;
  return traverse(head);
};
function traverse(right) {
  if (right == null) return true;
  let res = traverse(right.next);
  // 后序遍历
  res = res && right.val == left.val;
  left = left.next;
  return res;
}


// 思路2 : 双指针 + 反转链表的做法。 
// 整个流程可以分为以下五个步骤：
// 找到前半部分链表的尾节点。
// 反转后半部分链表。
// 判断是否回文。
// 恢复链表。
// 返回结果。
func reverseList(head *ListNode) *ListNode {
    var prev, cur *ListNode = nil, head
    for cur != nil {
        nextTmp := cur.Next
        cur.Next = prev
        prev = cur
        cur = nextTmp
    }
    return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
    fast := head
    slow := head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }

    // 找到前半部分链表的尾节点并反转后半部分链表
    firstHalfEnd := endOfFirstHalf(head)
    secondHalfStart := reverseList(firstHalfEnd.Next)

    // 判断是否回文
    p1 := head
    p2 := secondHalfStart
    result := true
    for result && p2 != nil {
        if p1.Val != p2.Val {
            result = false
        }
        p1 = p1.Next
        p2 = p2.Next
    }

    // 还原链表并返回结果
    firstHalfEnd.Next = reverseList(secondHalfStart)
    return result
}