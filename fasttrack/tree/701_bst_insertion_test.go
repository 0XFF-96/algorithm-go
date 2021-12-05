package tree


// 递归和迭代的解法
// 暂时不会


// 1、视频
// 2、
// BST insert: https://www.youtube.com/watch?v=bmaeYtlO2OE
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	var parent *TreeNode
	cur := root

	for cur != nil {
		parent = cur
		if cur.Val < val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	if parent.Val < val {
		parent.Right = &TreeNode{Val:val}
	} else {
		parent.Left = &TreeNode{Val:val}
	}
	return root
}

// 另一种方法
// https://leetcode.com/problems/insert-into-a-binary-search-tree/discuss/150757/java-iterative-100
// 递归的方法没有？


func insertIntoBSTV3(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return &TreeNode{Val: val}
	}


	if root.Val > val{
		root.Left = insertIntoBST(root.Left, val)
	}else{
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}


//public TreeNode insertIntoBST(TreeNode root, int val) {
//if(root == null) return new TreeNode(val);
//TreeNode cur = root;
//while(true) {
//if(cur.val <= val) {
//if(cur.right != null) cur = cur.right;
//else {
//cur.right = new TreeNode(val);
//break;
//}
//} else {
//if(cur.left != null) cur = cur.left;
//else {
//cur.left = new TreeNode(val);
//break;
//}
//}
//}
//return root;
//}