package tree



// https://www.youtube.com/results?search_query=Convert+BST+to+Greater+Tree
// youtube 视频



// 中序遍历
// [1, 2, 3]
// 记录 pre 的 value
// 左根右 -》 变成
// 然后累积 val
// 不用 morri 的方法
func convertBST(root *TreeNode) *TreeNode {
	if root ==  nil {
		return nil
	}

	stack := []*TreeNode{}
	cur := root
	var sumFromRightMost int
	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 这句
			// 区别
			val := cur.Val
			cur.Val += sumFromRightMost
			sumFromRightMost += val
			cur = cur.Left
			continue
		}

		stack = append(stack, cur)
		cur = cur.Right
	}
	return root
}


// 1、递归的解法
// 2、尝试用 Morii 的逆中序遍历解法
// 参考 https://leetcode.com/problems/convert-bst-to-greater-tree/solution/


// 递归解法，是中序遍历 递归形式之一
//