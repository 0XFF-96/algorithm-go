package tree

import (
	"testing"
)

func TestFlatten(t *testing.T){
	t1 := initTree()
	flatten(t1)
	//t.Log(t1.Right)
	//t.Log(t1.Left)
}


// 这是一个前序遍历的结果序列..
// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
// 这个前序遍历的结果应该没问题了吧

// 只是数值一样吗？
// 还是节点的存储地址也要一样?
func flatten(root *TreeNode){
	// 利用中序遍历的结果
	// succ
	// 中序遍历的框架

	curRoot := root
	stack := []*TreeNode{}
	// var parentRoot *TreeNode
	// dummyNode := &TreeNode{}
	// curDummy := dummyNode
	// var preLeftNode *TreeNode
	for len(stack)!=0 || curRoot != nil {
		if curRoot == nil {
			curRoot = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 链接起来
			// 重点语句有问题
			//curDummy.Right = curRoot
			//curDummy = curDummy.Right
			curRoot = curRoot.Right
			continue
		}

		// curDummy.Right = curRoot
		// curDummy = curDummy.Right

		// 在这里打印是层次遍历..
		// fmt.Println(curRoot.Val)
		stack = append(stack, curRoot)
		curRoot = curRoot.Left
	}
	// root = dummyNode.Right

}

// 递归解法, 牛
// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/discuss/37018/Recursive-Solution-in-Golang
// 实现遍历，但是不遍历叶子节点
// 从叶子节点的上一级判断...


//func Reverse(input string) string {
//	s := strings.Split(input, " ")
//	sort.Sort(sort.Reverse(sort.StringSlice(s)))
//	return strings.Join(s, " ")
//}

func TestFlattenV2(t *testing.T){
	tree1 := initTree()
	Flatten(tree1)
	t.Log(tree1)
	t.Log(tree1.Left)
	t.Log(tree1.Right)
	t.Log(tree1.Right.Right)

}

func Flatten(root *TreeNode){
	// 只取叶子节点的上一级节点
	if root == nil || (root.Left == nil && root.Right == nil){
		return
	}

	Flatten(root.Left)
	Flatten(root.Right)

	// 到了最左到叶子节点
	rootLeft := root.Left
	rootRight := root.Right
	root.Right = rootLeft
	root.Left = nil
	// cur := rootLeft
	for root.Right != nil {
		root = root.Right
	}
	root.Right = rootRight
}



// 非递归解法
// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/discuss/37125/I-think-there-is-something-wrong-with-Golang-test-environment