# 二叉树最大深度 

```
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0 
    }

    res := 0 
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        tmp := []*TreeNode{} // 有点浪费了， 其实用 pop 就可以了
        for _, node := range queue {
            if node.Left != nil {
                tmp = append(tmp, node.Left)
            }
            if node.Right != nil {
                tmp = append(tmp, node.Right)
            }
        }
        queue = tmp 
        res += 1 
    }
    return res 
}

```
