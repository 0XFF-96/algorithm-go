package tree

import (
	"fmt"
	"testing"
)


// 二次 AC
// 需要看看标准答案是怎么写的
//
func TestSumRootToLeaf(t *testing.T){
	node := &TreeNode{Val:1}
	node.Left = &TreeNode{Val:0}
	node.Right = &TreeNode{Val:1}
	node.Left.Left = &TreeNode{Val:0}
	node.Left.Right = &TreeNode{Val:1}

	node.Right.Left = &TreeNode{Val:0}
	node.Right.Right = &TreeNode{Val:1}

	ret := sumRootToLeaf(node)
	t.Log(ret)

}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var pathSum [][]int

	var help func(root *TreeNode, path []int)

	help = func(root *TreeNode, path []int) {
		if root == nil {
			return
		}

		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			// deep copy 问题
			pathSum = append(pathSum, path)
		}
		help(root.Left, path)
		help(root.Right, path)
		return
	}

	help(root, []int{})

	fmt.Println(pathSum)
	totalSum := 0

	for _, v := range pathSum {
		totalSum += countBit(v)
	}
	return totalSum
}

func countBit(b []int) int {
	sum := 0
	//for i:= len(b)-1;i > 0;i-- {
	//	//	if b[i] == 1 {
	//	//		fmt.Println(sum)
	//	//		sum += 1 << (i+1)
	//	//	}
	//	//}
	//for idx, v := range b {
	//	if v == 1 {
	//		sum += 1 << idx
	//	}
	//}

	ll := len(b)-1
	for idx, v := range b {
		sum += v << (ll -idx)
	}
	return sum
}

func TestCountBit(t *testing.T){
	r := countBit([]int{1, 0, 1})
	t.Log(r)
}


// 最巧妙之处在与
// bit 位的运算
// 来自于一个 google 工程师的解法
func sumRootToLeafV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return dfs(root, 0)
}

func dfs(root *TreeNode, cur int) int {
	if root == nil {
		return 0
	}

	// bit 运算
	cur = cur << 1 | root.Val

	if root.Left == nil && root.Right == nil {
		return cur
	}
	return dfs(root.Left, cur) + dfs(root.Right, cur)
}