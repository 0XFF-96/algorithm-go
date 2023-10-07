package tree

import (
	"container/list"
	"testing"
)

// construct file
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/


// 是否没有重复元素？
// 第一步 find the root
// 第二部 break down the tree
// 递归
// 能否用迭代的方法来做？

// 难点和时间消耗点在哪里？
// 如何进行递归？
// 如何进行非递归？

func initTree()*TreeNode{

	//   1
    //9     20
    //    15   7
var root = &TreeNode{Val:1}
root.Left = &TreeNode{Val:9}
root.Right = &TreeNode{Val:20}
root.Right.Right = &TreeNode{Val:7}
root.Right.Left = &TreeNode{Val:15}
return root
}

func TestBuildTree(t *testing.T){
	//root := initTree()

	t1 := buildTree([]int{1,2,3}, []int{2,1,3})
	t.Log(t1)
	t.Log(t1.Left.Val)
	t.Log(t1.Right.Val)

	//idx := []int{1, 2, 3}
	//t.Log(idx[:1])
}

func buildTree(preorder []int, inorder []int)*TreeNode{
	if len(preorder) != len(inorder){
		return nil
	}
	if len(preorder) < 2 || len(inorder)  < 2 {
		return &TreeNode{Val:preorder[0]}
	}
	root := &TreeNode{}
	root.Val = preorder[0]

	leftIndex := findRootIndex(preorder[0], inorder)

	// 边界条件
	// 特别注意⚠️
	// 这里是卡点地方
	// 原因是，没有了解到  slice[:] 的 左开右闭原则
	// 使用 leftIndex 不当
	// 解决的办法是
	// 利用小例子去，推测出索引应该怎么进行运算..
	//
	root.Left = buildTree(preorder[1:leftIndex+1], inorder[:leftIndex])
	root.Right = buildTree(preorder[leftIndex+1:], inorder[leftIndex+1:])

	// 1, 2, 3,   2, 1, 3
	return root
}


// 前提是一定有相同元素
func findRootIndex(val int, inorder[]int) int {
	for idx, i := range inorder {
		if i == val {
			return idx
		}
	}
	return 0
}

// 需要记录更多信息的迭代解法
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/discuss/34610/Golang-solution-(39ms)



// 迭代式的解法更加推崇
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/discuss/34545/Golang-iterative-solution-using-stack
// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/discuss/290526/O(n)-iterative-solution-in-golang
func buildTreePostOrder(inorder []int, postorder []int)*TreeNode{
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	stack := list.New()
	ii, pi := len(inorder)-1, len(postorder)-1
	root := &TreeNode{postorder[pi], nil, nil}
	pi--
	stack.PushBack(root)

	for pi >= 0 {
		topNode := stack.Back().Value.(*TreeNode)
		topMatch := false

		// 12345  12543
		for stack.Len() > 0 && stack.Back().Value.(*TreeNode).Val == inorder[ii] {
			// 从栈回溯，开始找左孩子
			topNode = stack.Remove(stack.Back()).(*TreeNode)
			ii--
			topMatch = true
		}

		// 将右孩子入栈，右孩子的关系
		node := &TreeNode{postorder[pi], nil, nil }
		pi--
		stack.PushBack(node)
		if topMatch{
			topNode.Left = node
		} else {
			topNode.Right = node
		}
	}

	return root
}

func buildTreeWithRec(inorder []int, postorder []int)*TreeNode{
	if len(inorder) == 0 {
		return nil
	}

	length := len(inorder)
	index := search(postorder[length-1], inorder)

	node := &TreeNode{
		Val:   postorder[length-1],
		Left:  buildTreeWithRec(inorder[:index], postorder[:index]),
		Right: buildTreeWithRec(inorder[index+1:], postorder[index:length-1]),
	}

	return node
}

func search(val int, array []int)int{
	for i := 0; i < len(array); i ++ {
		if val == array[i] {
			return i
		}
	}
	return 0
}


