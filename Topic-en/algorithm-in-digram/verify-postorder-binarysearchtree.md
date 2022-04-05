
# Verify Postorder binary search tree


```

func verifyPostorder(postorder []int) bool {
    // 左， 右， 根 。 ？ —— > 中序 是 有序的。 
    // 先恢复 - 》 再 判断？ 不行。 因为恢复是需要， 后序 和中序 。

    return recur(postorder, 0, len(postorder) - 1)
}

func recur(order []int, leftBound int,  rightBound int) bool {
    // base case ， 说明越界了
    if leftBound >= rightBound {
        return true 
    }

    leftIdx := leftBound
    
    // 这个感觉可以优化，不用每次递归都需要做这些事情？
    // 单调数组？ 需求是， 查询恰好大于此元素的边界。 
    // 
    for order[leftIdx] < order[rightBound] {
        leftIdx++
    }

    m := leftIdx

    for order[leftIdx] > order[rightBound] {
        leftIdx++
    }

    // 说明不匹配 
    // 根节点不唯一
    if leftIdx != rightBound {
        return false 
    }
    // recur 
    // 递归左子树， 递归右子树
    return recur(order, leftBound,  m -1 ) && recur(order, m, rightBound - 1)
}

```


