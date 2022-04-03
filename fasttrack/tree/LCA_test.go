package tree

// 这个是有暴力求解的思路
// 但是一下子，没有优雅的解决方案
//

// 这里有两个重要的前提
// All of the nodes' values will be unique.
// p and q are different and both values will exist in the BST.

// 搜索出两个序列
// 然后对比两个序列的不同节点..
// 就知道是不是了
// 有一种情况是, 父子节点的情况

// 伪代码实现思路
// 1、逻辑判断
// 2、实现一个函数，分别找到从根节点，到这两个节点的 []*TreeNode
// 3、对边这两个 slice, 返回分叉点 break

//

// 有一种情况不 ac ?
// 为什么？
//
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	leftPath := collectPath(root, p.Val)
	rightPath := collectPath(root, q.Val)

	ll := len(leftPath)
	lr := len(leftPath)

	var shortestPath []*TreeNode
	var longestPath []*TreeNode
	if ll > lr {
		shortestPath = leftPath
		longestPath = rightPath
	} else {
		shortestPath = rightPath
		longestPath = leftPath
	}

	var lcaNode *TreeNode
	for idx, n := range shortestPath {
		if n.Val != longestPath[idx].Val {
			// 第一次和最后一次的情况呢？
			return lcaNode
		}
		lcaNode = n
	}
	return lcaNode
}

func collectPath(root *TreeNode, val int) []*TreeNode {
	if root == nil {
		return nil
	}
	cur := root
	var ret []*TreeNode

	for cur != nil {
		ret = append(ret, cur)
		if cur.Val == val {
			return ret
		}
		// 有一个相等情况，怎么判断？
		if cur.Left != nil && cur.Left.Val > val {
			cur = cur.Right
		} else if cur.Left != nil && cur.Left.Val == val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return ret
}

// TODO: LCA
// LCA 的答案
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/
// 尝试用递归的方法去做
// 步骤
// 1、分解几种不同的情况 2、各种情况对应于什么操作？
// 3、找到上面的 code 不 ac 的原因 4、算法的伪代码描述

// lca 自己思考
// BST 的特点：1、不会有相等 2、左子树小于根小于右子树
// 寻找分叉点的方法.
// case 1: 落在同颗子树上，left or right
// case 2: 落在同一个子树，且是父子关系
// case 3: 不落在同一子树

// 算法伪代码
//

// 用递归解法
// 这颗子树都包含了目标节点 就是了

// LCA REFERENCE:
// 1、https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/
// 2、https://leetcode.com/articles/lowest-common-ancestor-of-a-binary-tree/
// 3、https://www.youtube.com/watch?v=TIoCCStdiFo // youtube 视频
//

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var helper func(root *TreeNode) bool
	var lca *TreeNode
	helper = func(root *TreeNode) bool {
		if root == nil {
			return false
		}

		left := helper(root.Left)

		right := helper(root.Right)

		// 中间变量
		// lol
		mid := root == p || root == q

		if mid && right || mid && left || right && left {
			lca = root
		}

		return mid || left || right
	}

	helper(root)
	return lca
}

// 视频里面的正解
// https://www.youtube.com/watch?v=TIoCCStdiFo
func lca(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if root.Val > max(p.Val, q.Val) {
		return lca(root.Right, p, q)
	} else if root.Val < min(p.Val, q.Val) {
		return lca(root.Left, p, q)
	} else {
		return root
	}
}

func lowestCommonAncestor22(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return root
	}
	// 寻找分叉点
	if root.Val > max(p.Val, q.Val) {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val < min(p.Val, q.Val) {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}

// 相关情况
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/discuss/293325/Golang-Solution-Iterative-Approach-24ms-beats-74.17-6.8MB-beats-82.30
// 好主意
func lcaFor(root, p, q *TreeNode) *TreeNode {
	val1, val2 := p.Val, q.Val
	cur := root

	for {
		switch {
		case val1 < cur.Val && val2 < cur.Val:
			cur = cur.Left
		case val1 > cur.Val && val2 > cur.Val:
			cur = cur.Right
		default:
			return cur
		}
	}
	//return root
}

// 如果不是一个 binary search tree 怎么办？
// https://www.youtube.com/watch?v=13m9ZCB8gjw&t=5s
//

// 有没有非递归做法？
//
func lowestCommonAncestorV22(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestorV22(root.Left, p, q)
	right := lowestCommonAncestorV22(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}

	if left == nil && right != nil {
		return right
	} else {
		return left
	}
}

// follow-up : 时间和空间复杂度是多少？
//
