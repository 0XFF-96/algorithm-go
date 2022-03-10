
- 递归。
- 简单迭代， 利用辅助栈 stack 。
- 修改指针指向的迭代，去除辅助 stack 的依赖。 



```

// 这个迭代有问题
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil 
    }
    result := []int{}
    
    // iterative approach 
    for root != nil {
        result = append(result, root.Val)
        
        if root.Left == nil {
            root = root.Right 
            continue 
        }
        
        if root.Right != nil {
            // move root.Right node into the end of right node of left 
            cur := root.Left // 
            for cur.Right != nil {
                cur = cur.Right
            }
            cur.Right = root.Right // ?
        }
        root = root.Left 
    }
    return result
}

```
