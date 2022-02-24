
- 简单滑动窗口值。
```

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }

    low, fast := 0, 0 
    for fast < len(nums) {
        if nums[low] != nums[fast] {
            low += 1
            nums[low] = nums[fast]
        }
        fast += 1 
    }
    return low + 1 
}

```

- 最长连续递增序列


```

func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }

    ret := 1
    for i := 0; i < len(nums); i ++ {
        cur := 1 
        for j := i + 1; j < len(nums); j ++ {
            if nums[j-1] < nums[j] {
                cur += 1 
            } else {
                break 
            }
        }
        ret = max(ret, cur)
    }
    return ret 
}

```

- 无重复字符的最长子串。
- 应该还有其他很多解法，需要整理一下。

```
func lengthOfLongestSubstring(s string) int {
    m := map[byte]int{}
    n := len(s)

    rk, ans := -1, 0 // 为什么初始化为这个？

    // 这道题目一定需要理解窗口的收缩方向是什么？
    // 窗口是如何进行收缩的？
    for i := 0; i < n; i ++ {
        if i != 0 {
            delete(m, s[i-1]) // 左指针移动，删除一个字符
        }
        for rk + 1 < n && m[s[rk + 1]] == 0 {
            // 不断移动右指针
            m[s[rk+1]] += 1
            rk++ 
        }

        // 第 i 到 rk 个字符是一个最长的无重复字符串。 
        ans = max(ans, rk - i + 1) // 为什么需要 减1 ？
    }
    return ans 
}
```

- 无重复字符串 V2 版本

```

func lengthOfLongestSubstring(s string) int {
    ans := 0 
    length := len(s)
    left, right := 0, 0 

    freqSet := map[byte]int{} // 判断是否重复元素。 map[byte]int{}, 不是  map[byte]bool{}
    for right < length { // right 不能等于 length， right = length -1  
        freqSet[s[right]] += 1
        // 如果重复，向左移动，直到没有重复元素出现
        for freqSet[s[right]] > 1 {
            freqSet[s[left]]-- 
            left++
        }

        ans = max(ans, right - left + 1)
        right++
    }

    return ans 
}

```

- 链表结合双指针

```

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    
    fast, slow := head, dummy 
    for i := 0; i < n; i++ {
        fast = fast.Next 
    }

    for ; fast != nil; fast = fast.Next {
        slow = slow.Next 
    }

    slow.Next = slow.Next.Next // 不会溢出，因为 fast 有相关的技巧了。
    return dummy.Next 
}

```

```

func middleNode(head *ListNode) *ListNode {
    slow := head 
    fast := head 

    for ; fast != nil &&fast.Next != nil; {
        slow = slow.Next 
        fast = fast.Next.Next 
    }
    return slow 
}

```

- 链表环检测2
- 重点在与此时的数学公式推导。

```

func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head 
    for fast != nil {
        slow = slow.Next 
        if fast.Next == nil {
            return nil // 没有环了
        }
        fast = fast.Next.Next 
        if fast == slow { // 相遇了
            // 重新开启计算
            p := head 
            for p != slow {
                p = p.Next
                slow = slow.Next
            }
            return p 
        }
    }
    return nil 
}

```

- 相交链表

```

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil 
    }

    pa, pb := headA, headB
    for pa != pb {
        if pa == nil { // 遍历完后，从 PB 开始
            pa = headB
        } else {
            pa = pa.Next 
        }

        if pb == nil { // 遍历完后，从 PA 开始
            pb = headA
        } else {
            pb = pb.Next 
        }
    }
    return pa 
}

```
