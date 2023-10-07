# k-largetst 
二叉搜索树的 k 大元素

- k 大元素， 所以需要倒序遍历。 这一点很不一样。 

```

func kthLargest(root *TreeNode, k int) int {
    var dfs func(root *TreeNode) 
    var ret int 

    // 全局 count k 
    countK := k 
    dfs = func(root *TreeNode) {
        if root == nil {
            return 
        }

        dfs(root.Right)
        if countK == 0 {
            return 
        }
        countK -= 1 
        if countK == 0 {
            ret = root.Val
        }

        dfs(root.Left)
    }
    dfs(root)
    return ret 
}

```
