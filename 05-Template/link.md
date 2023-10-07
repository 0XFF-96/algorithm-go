### 链表 

### 大纲


#### 递归反转链表
```
// 单链表节点的结构
public class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}
```

```
ListNode reverse(ListNode head) {
    if (head.next == null) return head;
    ListNode last = reverse(head.next); // 执行完后变怎么样？ // 执行完成后，整个链表就成了这样?
    head.next.next = head;
    head.next = null;
    return last;
}
不要跳进递归（你的脑袋能压几个栈呀？）
```

### reverseN + reverseBetween

```
ListNode reverseBetween(ListNode head, int m, int n) {
    // base case
    if (m == 1) {
        return reverseN(head, n);
    }
    // 前进到反转的起点触发 base case
    head.next = reverseBetween(head.next, m - 1, n - 1);
    return head;
}
```

### 回文

```
bool isPalindrome(string s) {
    int left = 0, right = s.length - 1;
    while (left < right) {
        if (s[left] != s[right])
            return false;
        left++; right--;
    }
    return true;
}
```

### 链表的 （前、后）序遍历方式

```
void traverse(ListNode head) {
    // 前序遍历代码
    traverse(head.next);
    // 后序遍历代码
}
```

type List struct {
    next *node 
}

type node struct {
    value int 
}

// 第一个

// 
// 3 - 5 - 6 - 2
// 5 - 3 


// 5- 6 保存 tmp 
// 5 - 3 
// 3 - tmp 



//
// 3 - 6 
// 5 - 3 
// 迭代
func reverseListV1(l *List) *List {
    if l == nil {
        return nil 
    }
    dummy := l 
    for dummy.next != nil {
        tmp := dummy.next.next 
        dummy.next = dummy
        next.next = tmp 
        dummy = dummy.next 
    }
    return dummy
}

// 递归
// 栈， 
func reverseListV2(head *List) *List {
    if head.next == nil {
        return head   // 
    }
    last = reverseListV2(l)
    head.next.next =  head  
    head.next = nil 
    return last
}



// 3 - 5 - 6 - 2

// 1. 
3 - reverseListV2([5, 6, 2 ])

// 
// 5 , 6 , 2 
// 6- 2
// 6 - 


### K 组翻转链表

- reverseKGroup 

```
    public ListNode reverseList(ListNode head) {
		ListNode pre = null, cur = head, next = null;
		while(cur!=null) {
			next = cur.next;
			cur.next = pre;
			pre = cur;
			cur = next;
		}
		return pre;
    }
```

```
    public ListNode reverseKGroup(ListNode head, int k) {
		// 建立一个哨兵节点用来返回
		ListNode dummy = new ListNode(0);
		dummy.next = head;
		
		// 两个指针用来保存需要反转数组的前后节点，用来恢复链表
		ListNode pre = dummy;
		ListNode next = dummy;
		
		// 需要一个 head 和 tail 用来保存被反转链表的头和尾
		ListNode tail = head;
		
		while(next!=null) {
			tail = head;
			// 找到尾节点
			for(int i=1;i<k && tail!=null;i++)
				tail = tail.next;
			
			// 如果不满足 k 个了，那么此时 tail 为 null，直接反转并跳出循环
			if(tail== null) {
				break;
			}
			
			//将尾结点的 next 保存并设为 null
			next = tail.next;
			tail.next = null;
			
			// 反转节点
			pre.next = reverse(head);
			
			// 重新赋值
			head.next = next;
			pre = head;
			head = head.next;
		}
		return dummy.next;
    }
```