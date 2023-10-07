# IS sub tree

- 一开始写出来和想出来来了， 但是判断 same tree 的时候，用了 ｜｜ 替代 && 。 导致一个测试用例有问题。 

- 但是为什么。。。 输出 true , 预期结果是 false 。 判断反了，方向有点松。

```

func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A == nil || B == nil {
        return false 
    }

    return  ( isSameTree(A, B) || 
    isSubStructure(A.Left, B) || 
    isSubStructure(A.Right, B))
}

func isSameTree(A *TreeNode, B *TreeNode) bool {
    if B == nil {
        return true 
    }
    if A == nil || A.Val != B.Val {
        return false 
    }

    return isSameTree(A.Left, B.Left) && isSameTree(A.Right, B.Right) // 这里是 && 才行。 // bug || 
}

```


