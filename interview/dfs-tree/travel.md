
- 树的遍历，非递归算法。 
- 前序。 

```

func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil 
    }
    
    stack := []*TreeNode{root}
    res := []int{}
    for len(stack) > 0 {
        node := stack[len(stack)-1]  
        stack = stack[:len(stack)-1] // pop 

        // 先右子树
        if node.Right != nil {
            stack = append(stack, node.Right)
        } 

        // 再左子树
        if node.Left != nil {
            stack = append(stack, node.Left)
        }

        res = append(res, node.Val)
    }
    return res 
}

```


- 中序遍历

```

func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    if root == nil {
        return result
    }

    stack := []*TreeNode{} // 初始化的时候，也需要注意，不需要提前放入 root . // why 
    for len(stack) != 0 || root != nil { // 两个结束条件
        for root != nil {
            stack = append(stack, root)
            root = root.Left // 一直往左深入。 左根右。
        }

        // 此时 root 已经是 nil, 代表没有 左子节点，开始出栈。
        root = stack[len(stack)-1] 
        stack = stack[:len(stack)-1] // pop 
        result = append(result, root.Val) 
        root = root.Right // 这个有可能为 nil root.Right == nil 
    }   
    return result
}

```
