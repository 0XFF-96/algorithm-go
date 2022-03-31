
# IS Balanced Tree 

```

func isBalanced(root *TreeNode) bool {
    // 1. 二叉树高度。
    // 2. 判断节点是否平衡。 
    // 3. 后序遍历验证
    if root == nil {
        return true 
    }

    // 每次开启递归，这部分都需要重复计算
    // 主要消耗在这里。 
    if abs(depth(root.Left) - depth(root.Right)) > 1 {
        return false 
    }

    // 开始递归
    // 这一步会产生重复计算。 
    return isBalanced(root.Left) && isBalanced(root.Right)
}


func depth(root *TreeNode) int {
    if root == nil {
        return 0 
    }


    return max(depth(root.Left), depth(root.Right)) + 1 
}

func abs(i int) int {
    if i < 0 {
        return -i 
    } else {
        return i 
    }
}

func max(i, j int) int {
    if i > j {
        return i 
    } else {
        return j 
    }
}

```
