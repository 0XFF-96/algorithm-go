package tree



// https://leetcode.com/problems/diameter-of-binary-tree/


// 递归的方式解法
//
// 两个难点
// 1、全局最大路径的记录？
// 2、递归返回值是什么？
// 3、
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 1
	}
	ans := 1
	depth(root, &ans)
	return ans -1
}

func depth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left, ans)
	r := depth(root.Right, ans)

	*ans = max(*ans, l+r+1)
	return max(l, r) + 1

}

// https://leetcode.com/problems/diameter-of-binary-tree/solution/
// 答案

// 有没有非递归解法？
//
