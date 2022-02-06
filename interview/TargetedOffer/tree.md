
- Tree  

```
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil 
    }

    queue := []*TreeNode{root}
    ret := make([][]int, 0)

    isLeft := true  
    for len(queue) != 0 {
        size := len(queue) 
        levelVals := make([]int, 0)
        for i := 0; i < size; i ++ {
            // dequeue
            treeNode := queue[0]
            queue = queue[1:]

            levelVals = append(levelVals, treeNode.Val)
            if treeNode.Left != nil {
                queue = append(queue, treeNode.Left)
            }
            if treeNode.Right != nil {
                queue = append(queue, treeNode.Right)
            }
        }
        if !isLeft {
                levelVals = reverse(levelVals)
            }
        isLeft = !isLeft
        ret = append(ret, levelVals)
    }

    return ret 
}

func reverse(s []int) []int {
    for i:=0; i < len(s)/2 ; i ++ {
        s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
    }
    return s 
}

```

- LCA ， 最近祖先


```
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if root == nil {
      return nil 
  }

  // 这种情况是
  //    
  //      3
  //    5
  //  1  2 
  // p = 5, q = 1 时，需要保证返回 5 。 是最小公共节点
  // 如何层层向上传递信息。 


  if root == p || root == q {
      return root 
  }

  // 开始搜索
  left := lowestCommonAncestor(root.Left, p , q)
  right := lowestCommonAncestor(root.Right,p, q)

  // 到达叶子节点
  if left == nil && right == nil {
      return nil 
  } else if left == nil && right != nil {
      return right
  } else if left != nil && right == nil {
      return left 
  } else {
      // left != nil && right != nil 
      return root 
  }
}
```
