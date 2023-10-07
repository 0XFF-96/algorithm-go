

# 反转二叉树🌲 


```

func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil 
    }

    // flip 
    // 前序遍历， 递归
    tmp := root.Right 
    root.Right = root.Left
    root.Left = tmp 
    if root.Left != nil {
        mirrorTree(root.Left)
    }
    if root.Right != nil {
        mirrorTree(root.Right)
    }
    return root
}

```


```

func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil 
    }
    stack := []*TreeNode{root}
    // flip 
    // 前序遍历， 递归

    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1] // 这里不能用 : , 否则会创建新变量，len(stack) > 0 的条件就不准了

        if node.Left != nil {
            stack = append(stack, node.Left)
        }

        if node.Right != nil {
            stack = append(stack, node.Right)
        }

        // 交换
        node.Left, node.Right = node.Right, node.Left
    }

    return root
}

```


- 是否镜像二叉树 

- 错误版本

```

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    // 后序遍历
    isLeft := isSymmetric(root.Left)
    isRight := isSymmetric(root.Right)

    ok := false 
    if root.Right != nil && root.Left != nil && root.Right.Val == root.Left.Val {
        ok = true 
    }
    if root.Right == nil && root.Left == nil {
        ok = true
    }

    return isLeft && isRight && ok 
}

```


```

func isSymmetric(root *TreeNode) bool {
    return root == nil || recur(root.Left, root.Right)
}


func recur(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    // 提前剪枝
    if left == nil || right == nil || left.Val != right.Val {
        return false 
    }

    // 递归深入, 注意入参数
    return recur(left.Left, right.Right) && recur(left.Right, right.Left)
}

```
