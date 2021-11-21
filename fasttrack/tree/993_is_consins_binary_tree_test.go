package tree




// 一次 AC
// 爽
// 难点和重点都在与判断
// 如何判断是否 拥有同一个父亲节点上


// ===========
// 可以用 level order 迭代解法？
// https://leetcode.com/problems/cousins-in-binary-tree/discuss/240992/Golang-BFS
// solution 2
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	// depth := 0
	var preorder func(root *TreeNode, depth int, pos int)

	var xNode *AnodeV2
	var yNode *AnodeV2

	preorder = func(root *TreeNode, depth int, pos int) {
		if root == nil {
			return
		}
		depth++
		if root.Val == x || root.Val == y {
			if xNode == nil {
				xNode = &AnodeV2{node: root, depth:depth, pos:pos}
			} else {
				yNode = &AnodeV2{node: root, depth:depth, pos:pos}
			}
		}
		preorder(root.Left, depth, pos*2)
		preorder(root.Right, depth, pos*2+1)
	}
	preorder(root, 0, 1)

	if xNode == nil || yNode == nil {
		return false
	}

	if xNode.depth != yNode.depth {
		return false
	}


	// 这个判断条件有问题
	//
	//     xPPos := xNode.pos / 2
	//     yPPos := (yNode.pos -1) / 2

	//     if xPPos == yPPos {
	//         return false
	//     }

	if xNode.pos +1 == yNode.pos {
		return false
	}
	return true
}

type AnodeV2 struct {
	node *TreeNode
	depth int
	pos int
}

