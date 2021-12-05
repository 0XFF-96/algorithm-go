package tree

import (
	"fmt"
	"strconv"
	"testing"
)

// 问题的思路
// 伪代码描述

// 答案：https://leetcode.com/problems/print-binary-tree/solution/
// 难点，树节点之间与最终的 array 有什么关联？
// 如何确定最终 index 不出错？
//
// https://leetcode.com/problems/print-binary-tree/discuss/123034/golang-BFS-solution
// 非递归解法
//
// follow-up-Question: 你能否想出非递归解法
func TestPrintTree(t *testing.T){
	// tree1 := initTree()
	// mathutil.Max()

	tree1 := initTree()
	array := printTree(tree1)

	t.Log(array)
}

func printTree(root *TreeNode) [][]string {
	height := getHeight(root)
	col := (1 << height) - 1
	row := height

	fmt.Println(height, col)
	ret := make([][]string, row)
	for i :=0;i < row; i++ {
		ret[i] = make([]string, col)
		for j:=0;j<col; j++ {
			ret[i][j] = ""
		}
	}
	fill(ret, root, 0, 0, col)
	return ret
}

func fill(res [][]string, root *TreeNode, i, l, r int){
	if root == nil {
		return
	}
	res[i][(l+r)/2] =  strconv.Itoa(root.Val)
	fill(res, root.Left, i+1, l, (l+r)/2)
	fill(res, root.Right, i +1, (l+r+1)/2, r)
}



func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.Left),getHeight(root.Right))
}