package tree

// morii  中序遍历做一下
// 一次 AC
// 中序遍历 Morri 算法
//
func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	cur := root

	var pre *TreeNode
	for cur != nil {
		if cur.Left == nil {
			if cur.Val>=L && cur.Val <= R {
				sum += cur.Val
			}
			cur = cur.Right
			continue
		} else {
			pre = cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else {
				pre.Right = nil
				if cur.Val>=L && cur.Val <= R {
					sum += cur.Val
				}
				cur = cur.Right
			}
		}
	}
	return sum
}


