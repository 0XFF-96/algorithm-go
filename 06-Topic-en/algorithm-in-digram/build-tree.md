# 重建二叉树

1. 从中序和后序序列，重建二叉树🌲 。 

- 这种写法不能扩展，比较恶心。🤢 。 

```

func buildTree(preorder []int, inorder []int) *TreeNode {
    m := map[int]int{}
    for i := 0; i < len(inorder); i ++ {
        m[inorder[i]] = i // 快速找到 对应的索引 
    }

    var recur func(root int, left int, right int) *TreeNode

    recur = func(root int, left int ,  right int) *TreeNode {
        if left > right {
            return nil 
        }
        node := &TreeNode{Val: preorder[root]}

        i := m[preorder[root]]

        node.Left = recur(root + 1, left, i -1) // 左子树递归
        node.Right = recur(root + i - left + 1, i + 1, right) // 右子树递归
        return node // 返回根节点 
    }

    return recur(0, 0, len(inorder) -1)
}

```
