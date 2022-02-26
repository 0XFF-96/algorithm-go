
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

- 后序遍历（写法一，利用头插法的特点） （但是最好不要用） （用一种和上面写法一样的方法）

```

func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }

    stack := []*TreeNode{root}
    for len(stack) != 0 {
        curr := stack[len(stack) -1] 
        stack = stack[:len(stack)-1]

        if curr.Left != nil {
            stack = append(stack, curr.Left)
        }

        if curr.Right != nil {
            stack = append(stack, curr.Right)
        }
        // 后序， 左右根。
        res = append(res, curr.Val) // 感觉不太对。 // addFirst 的问题，前插法。最好不要用这种。
    }
    // return res 
    return reverseArray(res) 
}

func reverseArray(a []int) []int{
    // return a 
    l := len(a) - 1 
    for i :=0; i <= ( len(a) -1 )/ 2; i ++ {
        tmp := a[l-i]
        a[l-i] = a[i]
        a[i] = tmp  
        // a[i], a[l - i] = a[l-1], a[i]
    }
    return a 
}

```

- 多叉树的后序遍历方法。 

```

func postorder(root *Node) []int {
    if root == nil {
        return nil 
    }
    var ret []int 
    for _, n := range root.Children {
        r := postorder(n)
        ret = append(ret, r...)
    }
    ret = append(ret, root.Val)
    return ret 
}

```
