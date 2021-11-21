package tree

// 从不同的节点
// 从祖节点的角度出发。 左子树 == 右子树


// 学到了一个技巧，可以将数据复制多一份，然后进行遍历
// https://leetcode.com/problems/symmetric-tree/discuss/495307/Go-iterative-with-two-stacks-based-queue-and-recursive-solutions
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root, root)
}

func isSymmetricHelper(t1, t2 *TreeNode)bool {
	if (t1 == nil && t2 != nil ) ||
		(t1 != nil && t2 != nil ){
		return false
	}
	if t1.Val != t2.Val {
		return false
	}

	return isSymmetricHelper(t1.Left, t2.Right) && isSymmetricHelper(t1.Right, t2.Left)
}


// TODO:版本
// 非递归版本：明天再看https://leetcode.com/problems/symmetric-tree/discuss/495307/Go-iterative-with-two-stacks-based-queue-and-recursive-solutions
func isSymeetricWithStack(root *TreeNode) bool{

	return false
}


// 2 月 9 号
//
func isSymmetricV2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 递归的解法
	// 迭代的呢？
	return isEqual(root.Left, root.Right)
}

func isEqual(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}

	left := isEqual(r1.Left, r2.Right)
	right := isEqual(r1.Right, r2.Left)

	return left && right && r1.Val == r2.Val

}


