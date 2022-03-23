### 二叉树到层次遍历

```


func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil 
    }

    var res [][]int 
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        tmp := make([]int, 0, len(queue))
        
        for i := len(queue) ; i > 0; i -- {
            node := queue[0]
            queue = queue[1:]
            tmp = append(tmp, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }

        }

        res = append(res, tmp)
    }
    return res 
}

```
