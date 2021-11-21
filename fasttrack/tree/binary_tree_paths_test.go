package tree

import (
	"strconv"
	"strings"
	"testing"
)

// 这一题的解决思路和 sum number 非常类似
// 只需要改造一下其中的几行代码即可...


func TestBinaryTreePaths(t *testing.T){
	tree1 := initTree()

	strPath := binaryTreePaths(tree1)
	t.Log(strPath)
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	// 这条路径有没有返回次序之分？
	var ret []string
	var helper = func(root *TreeNode, path []string){}
	helper = func(root *TreeNode, path []string){
		if root == nil {
			return
		}

		// ----
		path = append(path, strconv.Itoa(root.Val))
		helper(root.Left, path)
		helper(root.Right, path)
		// --

		// 这几行代码的放置顺序非常重要
		if root.Left == nil && root.Right == nil {
			// fmt.Printf("%v", path)
			str := strings.Join(path, "->")
			ret = append(ret, str)
		}
		return
	}
	helper(root, []string{})

	return ret
}