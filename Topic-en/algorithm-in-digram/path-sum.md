### 二叉树路径和 


```

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
    var res [][]int 
    var path []int 

    var recur func(root *TreeNode, target int) // 需要先定义，后才能在函数内部使用
    recur = func(root *TreeNode, target int) {
        if root == nil {
            return 
        }

        path = append(path, root.Val)
        target = target - root.Val
        if target == 0 && root.Left == nil && root.Right == nil {
            curPath := append([]int{}, path...)
            res = append(res, curPath)
        }
        recur(root.Left, target)
        recur(root.Right, target)
        path = path[:len(path) -1]
    }

    recur(root, target)
    return res
}


```
