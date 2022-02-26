
- 二叉树的堂兄弟节点

```

https://leetcode-cn.com/problems/cousins-in-binary-tree/solution/


func isCousins(root *TreeNode, x int, y int) bool {
    if root == nil {
        return false
    }
    q := new(Queue)
    q.EnQueue(root)
    for q.Size() > 0 {
        n := q.Size()
        p := make(map[int]*TreeNode)
        for i := 0; i < n; i++ {
            node := q.DeQueue()
            if node.Left != nil {
                q.EnQueue(node.Left)
                if node.Left.Val == x || node.Left.Val == y {
                    p[node.Left.Val] = node
                }
            }
            if node.Right != nil {
                q.EnQueue(node.Right)
                if node.Right.Val == x || node.Right.Val == y {
                    p[node.Right.Val] = node
                }
            }
        }
        if len(p) == 2 && p[x] != nil && p[y] != nil && p[x] != p[y] {
            return true
        }
    }
    return false
}

type Queue struct {
    queue []*TreeNode
}

func (q Queue)Size() int {
    return len(q.queue)
}

func (q *Queue)EnQueue(node *TreeNode) {
    if q.queue == nil {
        q.queue = make([]*TreeNode, 0)
    }
    q.queue = append(q.queue, node)
}

func (q *Queue)DeQueue() *TreeNode {
    if q.Size() == 0 {
        return nil
    }
    node := q.queue[0]
    q.queue = q.queue[1:]
    return node
}

```
