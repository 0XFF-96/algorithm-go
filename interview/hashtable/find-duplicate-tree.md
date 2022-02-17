- 寻找相同子树。 

```

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    res := []*TreeNode{}
    if root == nil {
        return res 
    }

    subTreeMap := make(map[string]int)
    var dfs func(*TreeNode) string 

    dfs = func(node *TreeNode) string {
        if node == nil {
            return "#"
        }

        subTree := fmt.Sprintf("%v:%v:%v", node.Val,
            dfs(node.Left),
            dfs(node.Right),
        )
        subTreeMap[subTree] += 1 

        if subTreeMap[subTree] == 2 {
            res = append(res, node)
        }
        return subTree
    }
    dfs(root)
    return res 
}

```
