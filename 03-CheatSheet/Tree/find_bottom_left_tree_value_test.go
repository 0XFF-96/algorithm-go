package tree


// https://leetcode.com/problems/find-bottom-left-tree-value/
// https://leetcode.com/problems/find-bottom-left-tree-value/discuss/98779/Right-to-Left-BFS-(Python-%2B-Java)

// 需要讨论几个 test case 不过的情况下
// https://leetcode.com/problems/find-bottom-left-tree-value/discuss/98779/Right-to-Left-BFS-(Python-%2B-Java)
// https://www.youtube.com/watch?v=5Tkw1NRJ5Ec
// 层次遍历的相关解法：
// https://www.youtube.com/watch?v=7uG0gLDbhsI
// not just that...


// 用 level order 的思想
// 但是转换了一下位置
// 层次遍历的递归解法...
//

// 难点在哪里？
// 有可能的 follow-up 的 question
//
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	var leftmost int
	for len(queue) != 0 {
		size := len(queue)
		for i:=0;i<size;i++{
			top := queue[0]
			queue = queue[1:]
			if i == size -1 {
				leftmost = top.Val
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
		}
	}

	return leftmost
}


// 尝试用递归的方法解法
// https://leetcode.com/problems/find-bottom-left-tree-value/discuss/98806/C%2B%2B-Clean-Code-DFS-Recursion-with-Explanation
//

// BFS + queue
// DFS + queue
// https://leetcode.com/problems/find-bottom-left-tree-value/discuss/98770/Python-DFS-%2B-Stack-and-BFS-%2B-Queue-Solution
