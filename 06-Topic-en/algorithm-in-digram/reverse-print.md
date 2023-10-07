
- 从尾部打印链表
- 递归打印 【数组】 ？

```

func reversePrint(head *ListNode) []int {
    if head == nil {
        return nil
    }

    res := reversePrint(head.Next)
    res = append(res, head.Val)
    return res 
}
```

- 这道题不难，倒是重新复习了一遍 求中点点各个解法。 
- 坑在于 reverseArray 的几种求解方法。 
- 其实，最清晰一点的就是 最后面一种了，为什么。 不用 length - i - 1 。 （双指针，一个方向交替移动）


```

func reversePrint(head *ListNode) []int {
    tmpArray := []int{}
    for head != nil {
        tmpArray = append(tmpArray, head.Val)
        head = head.Next 
    }

    return reverseArray(tmpArray)
}

func reverseArray(array []int) []int {
    lastIdex := len(array) - 1 // last index 
    // length := len(array) // 不要这样子，清晰一点的写法是

    // 在这里踩坑了。
    // i < (len(array)/ 2 ) 
    // i <= (len(array) / 2 - 1)
    for i := 0; i < ( len(array)/ 2) ; i++ { 
        array[i], array[ lastIdex - i ] = array[lastIdex - i], array[i]
    }
    return array 
}


// func reverseArrayV2(array []int) []int {
//     // lastIdex := len(array) - 1 // last index 
//     // length := len(array) // 不要这样子，清晰一点的写法是

//     // 在这里踩坑了。
//     // i < (len(array)/ 2 ) 
//     // i <= (len(array) / 2 - 1)
//     for i := 0; i < ( length / 2) ; i++ { 
//         array[i], array[ length - i -1] = array[length - i -1], array[i]
//     }
//     return array 
// }

```
