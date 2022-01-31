

// 对称二叉树
// 镜像对称二叉树


// 你可以运用递归和迭代两种方法解决这个问题吗？
// 递归很简单， 需要提前

// 递归
// 对于此题： 递归的点怎么找？从拿到题的第一时间开始，思路如下：
// 对于此题： 递归的点怎么找？从拿到题的第一时间开始，思路如下：

// 1.怎么判断一棵树是不是对称二叉树？ 答案：如果所给根节点，为空，那么是对称。如果不为空的话，当他的左子树与右子树对称时，他对称

// 2.那么怎么知道左子树与右子树对不对称呢？在这我直接叫为左树和右树 答案：如果左树的左孩子与右树的右孩子对称，左树的右孩子与右树的左孩子对称，那么这个左树和右树就对称。

// 仔细读这句话，是不是有点绕？怎么感觉有一个功能A我想实现，但我去实现A的时候又要用到A实现后的功能呢？

// 当你思考到这里的时候，递归点已经出现了： 递归点：我在尝试判断左树与右树对称的条件时，发现其跟两树的孩子的对称情况有关系。

// 想到这里，你不必有太多疑问，上手去按思路写代码，函数A（左树，右树）功能是返回是否对称

// def 函数A（左树，右树）： 左树节点值等于右树节点值 且 函数A（左树的左子树，右树的右子树），函数A（左树的右子树，右树的左子树）均为真 才返回真


func isSymmetric(root *TreeNode) bool {
    return check(root, root)
}

func check(p, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left) 
}


// 迭代

// 层次遍历
// 而且每一层的一个数组，是回文数组！
// 这个思路还有一些缺陷？

// 「方法一」中我们用递归的方法实现了对称性的判断，那么如何用迭代的方法实现呢？
// 首先我们引入一个队列，这是把递归程序改写成迭代程序的常用方法。初始化时我们把根节点入队两次。
// 每次提取两个结点并比较它们的值（队列中每两个连续的结点应该是相等的，而且它们的子树互为镜像），
// 然后将两个结点的左右子结点按相反的顺序插入队列中。当队列为空时，
// 或者我们检测到树不对称（即从队列中取出两个不相等的连续结点）时，该算法结束。


// 大致思路：
// 使用BFS（迭代）
// 一次向队列中加入两个节点，这样出队的时候也一次出两个节点
// 这两个节点进行比较即可，然后将他们各自的左右节点入队
// 需要注意的是，应该将镜像位置的节点成对入队，不然就错了，
// 因为是一次性是处理两个节点


func isSymmetric(root *TreeNode) bool {
    u, v := root, root
    q := []*TreeNode{}
    q = append(q, u)
    q = append(q, v)
    for len(q) > 0 {
        u, v = q[0], q[1]
        q = q[2:]
        if u == nil && v == nil {
            continue
        }
        if u == nil || v == nil {
            return false
        }
        if u.Val != v.Val {
            return false
        }
        q = append(q, u.Left)
        q = append(q, v.Right)

        q = append(q, u.Right)
        q = append(q, v.Left)
    }
    return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			left := queue[0]
			right := queue[1]
			queue = queue[2:]
			size = size - 2
			if left == nil && right == nil {
				continue
			}
			if left == nil || right == nil || left.Val != right.Val {
				return false
			}
			queue = append(queue, left.Left)
			queue = append(queue, right.Right)
			queue = append(queue, left.Right)
			queue = append(queue, right.Left)
		}
	}
	return true
}
